package mapservice

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/RyuChk/ips-map-service/internal/constants"
	"github.com/RyuChk/ips-map-service/internal/models"
	"github.com/RyuChk/ips-map-service/internal/repository/minio"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	buildingcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/buildingCollectionRepo"
	floorcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/floorCollectionRepo"
	floordetailcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/floorDetailCollectionRepo"
	rediscache "github.com/RyuChk/ips-map-service/internal/repository/redisCache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source=service.go -destination=mock_mapService/mock_service.go -package=mock_mapservice
type Service interface {
	GetFloorListByBuilding(ctx context.Context, building string, role constants.UserRole) ([]models.Floor, error)
	GetFloorDetailFromKey(ctx context.Context, floor int, building string, role constants.UserRole) (models.FloorDetail, error)
	AddBuildingInfoToDB(ctx context.Context, body models.Building) error
	AddFloorToDB(ctx context.Context, body models.Floor) error
	UpsertFloorDetail(ctx context.Context, body models.FloorDetail) error
	GetBuildingInfo(ctx context.Context, building string, role constants.UserRole) (models.Building, error)
	GetBuildingList(ctx context.Context, admin bool) ([]models.Building, error)
}

type service struct {
	minio                     minio.Service
	buildingCollectionRepo    buildingcollectionrepo.Repository
	floorCollectionRepo       floorcollectionrepo.Repository
	floorDetailCollectionRepo floordetailcollectionrepo.Repository
	redisClient               rediscache.Service
}

func ProvideMapURLService(minio minio.Service, buildingCollectionRepo buildingcollectionrepo.Repository, floorCollectionRepo floorcollectionrepo.Repository, floorDetailCollectionRepo floordetailcollectionrepo.Repository, rediscache rediscache.Service) Service {
	return &service{
		minio:                     minio,
		buildingCollectionRepo:    buildingCollectionRepo,
		floorCollectionRepo:       floorCollectionRepo,
		floorDetailCollectionRepo: floorDetailCollectionRepo,
		redisClient:               rediscache,
	}
}

func (s *service) GetFloor(ctx context.Context, building string, floor int) (models.Floor, error) {
	floorCachePrefix := fmt.Sprintf("%s%s:%d", constants.FloorCachePrefix, building, floor)
	raw, err := s.redisClient.Get(ctx, floorCachePrefix)
	if err != nil {
		filter, _ := mongodb.AddFilter(mongodb.Filter{"building": building}, mongodb.Filter{"floor": floor})
		result, err := s.floorCollectionRepo.FindOne(ctx, filter)
		if err != nil {
			return models.Floor{}, status.Error(codes.Internal, err.Error())
		}

		json, err := json.Marshal(result)
		if err != nil {
			return models.Floor{}, status.Error(codes.Internal, err.Error())
		}

		var value interface{} = json
		s.redisClient.Set(ctx, floorCachePrefix, value, time.Hour)
		return result, nil
	}

	result := models.Floor{}
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return models.Floor{}, status.Error(codes.Internal, err.Error())
	}

	return result, nil
}

func (s *service) GetBuildingList(ctx context.Context, admin bool) ([]models.Building, error) {
	cachePrefix := constants.BuildingCachePrefix.String() + "LIST"

	raw, err := s.redisClient.Get(ctx, cachePrefix)
	if err != nil {
		result, err := s.buildingCollectionRepo.Find(ctx, mongodb.Filter{})
		if err != nil {
			return []models.Building{}, status.Error(codes.Internal, err.Error())
		}

		json, err := json.Marshal(result)
		if err != nil {
			return []models.Building{}, status.Error(codes.Internal, err.Error())
		}

		var value interface{} = json
		s.redisClient.Set(ctx, cachePrefix, value, time.Hour)
		return result, nil
	}

	result := []models.Building{}
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return []models.Building{}, status.Error(codes.Internal, err.Error())
	}

	return result, nil
}

func (s *service) GetBuildingInfo(ctx context.Context, building string, role constants.UserRole) (models.Building, error) {
	result := models.Building{}

	buildingCachePrefix := constants.BuildingCachePrefix.String() + building
	raw, err := s.redisClient.Get(ctx, buildingCachePrefix)
	if err != nil {
		result, err := s.buildingCollectionRepo.FindOne(ctx, mongodb.Filter{"name": building})
		if err != nil {
			return models.Building{}, status.Error(codes.Internal, err.Error())
		}

		json, err := json.Marshal(result)
		if err != nil {
			return models.Building{}, status.Error(codes.Internal, err.Error())
		}

		var value interface{} = json
		s.redisClient.Set(ctx, buildingCachePrefix, value, time.Hour)

		floorList, err := s.GetFloorListByBuilding(ctx, building, role)
		if err != nil {
			return models.Building{}, status.Error(codes.Internal, err.Error())
		}

		result.FloorList = floorList
		return result, nil
	}

	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return models.Building{}, status.Error(codes.Internal, err.Error())
	}

	floorList, err := s.GetFloorListByBuilding(ctx, building, role)
	if err != nil {
		return models.Building{}, status.Error(codes.Internal, err.Error())
	}

	result.FloorList = floorList
	return result, nil
}

func (s *service) GetFloorListByBuilding(ctx context.Context, building string, role constants.UserRole) ([]models.Floor, error) {
	floorListCachePrefix := fmt.Sprintf("%sLIST:%s-%s", constants.FloorCachePrefix, building, role)

	raw, err := s.redisClient.Get(ctx, floorListCachePrefix)
	if err != nil {
		filter, _ := mongodb.AddFilter(mongodb.Filter{"building": building}, mongodb.Filter{"is_admin": false})
		if role == constants.AdminRole {
			filter = mongodb.Filter{"building": building}
		}

		floorList, err := s.floorCollectionRepo.Find(ctx, filter)
		if err != nil {
			return []models.Floor{}, status.Error(codes.Internal, err.Error())
		}

		json, err := json.Marshal(floorList)
		if err != nil {
			return []models.Floor{}, status.Error(codes.Internal, err.Error())
		}

		var value interface{} = json
		s.redisClient.Set(ctx, floorListCachePrefix, value, time.Hour)
		return floorList, nil
	}

	result := []models.Floor{}
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return []models.Floor{}, status.Error(codes.Internal, err.Error())
	}

	return result, nil
}

func (s *service) GetFloorDetailFromKey(ctx context.Context, floor int, building string, role constants.UserRole) (models.FloorDetail, error) {
	_floor, err := s.GetFloor(ctx, building, floor)
	if err != nil {
		return models.FloorDetail{}, status.Error(codes.NotFound, "cannot found floor in database make sure your input is correct")
	}

	if _floor.IsAdmin && role != constants.AdminRole {
		return models.FloorDetail{}, status.Error(codes.PermissionDenied, "your do not have permission to view this data")
	}

	filter := mongodb.Filter{"key": fmt.Sprintf("%s-%d", building, floor)}
	floorDetail, err := s.floorDetailCollectionRepo.FindOne(ctx, filter)
	if err != nil {
		return models.FloorDetail{}, status.Error(codes.NotFound, "cannot found map detail in database make sure your input is correct")
	}

	floorDetail.Info = _floor
	return floorDetail, nil
}

func (s *service) AddBuildingInfoToDB(ctx context.Context, body models.Building) error {
	filter := mongodb.Filter{"building": body.Name}
	_, err := s.buildingCollectionRepo.FindOne(ctx, filter)
	if err == nil {
		return status.Error(codes.AlreadyExists, "building information is already exist in database")
	}

	if err := s.buildingCollectionRepo.InsertOne(ctx, body); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *service) AddFloorToDB(ctx context.Context, body models.Floor) error {
	_, err := s.GetFloor(ctx, body.Building, body.Floor)
	if err == nil {
		return status.Error(codes.AlreadyExists, "floor information is already exist in database")
	}

	if err := s.floorCollectionRepo.InsertOne(ctx, body); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *service) UpsertFloorDetail(ctx context.Context, body models.FloorDetail) error {
	if err := s.floorDetailCollectionRepo.Upsert(ctx, body); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
