package mapservice

import (
	"context"
	"fmt"
	"time"

	"github.com/RyuChk/ips-map-service/internal/constants"
	"github.com/RyuChk/ips-map-service/internal/models"
	"github.com/RyuChk/ips-map-service/internal/repository/minio"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	mapcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/mapCollectionRepo"
	mapurlcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/mapURLCollectionRepo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source=service.go -destination=mock_mapService/mock_service.go -package=mock_mapservice
type Service interface {
	GetFloorListByBuilding(ctx context.Context, building string, role constants.UserRole) ([]models.Map, error)
	AddFloorToDB(ctx context.Context, body models.Map) error
	AddMapURL(ctx context.Context, floor int, building, url string) (models.MapImageURL, error)
	GetMapURLFromKey(ctx context.Context, floor int, building string, role constants.UserRole) (models.MapImageURL, error)
}

type service struct {
	minio                minio.Service
	mapURLCollectionRepo mapurlcollectionrepo.Repository
	mapCollectionRepo    mapcollectionrepo.Repository
}

func ProvideMapURLService(minio minio.Service, mapURLCollectionRepo mapurlcollectionrepo.Repository, mapCollectionRepo mapcollectionrepo.Repository) Service {
	return &service{
		minio:                minio,
		mapURLCollectionRepo: mapURLCollectionRepo,
		mapCollectionRepo:    mapCollectionRepo,
	}
}

func (s *service) GetFloorListByBuilding(ctx context.Context, building string, role constants.UserRole) ([]models.Map, error) {
	filter := mongodb.Filter{"building": building}
	if role != constants.AdminRole {
		f, _ := mongodb.AddFilter(filter, mongodb.Filter{"is_admin": false})
		filter = f
	}
	floors, err := s.mapCollectionRepo.Find(ctx, filter)
	if err != nil {
		return []models.Map{}, status.Error(codes.NotFound, err.Error())
	}

	return floors, nil
}

func (s *service) AddMapURL(ctx context.Context, floor int, building, url string) (models.MapImageURL, error) {
	checkExistFilter := mongodb.Filter{"key": fmt.Sprintf("%s-%d", building, floor)}
	_, err := s.mapURLCollectionRepo.FindOne(ctx, checkExistFilter)
	if err == nil {
		return models.MapImageURL{}, status.Error(codes.AlreadyExists, "map url information is already exist in database")
	}

	filter, _ := mongodb.AddFilter(mongodb.Filter{"building": building}, mongodb.Filter{"number": floor})
	m, err := s.mapCollectionRepo.FindOne(ctx, filter)
	if err != nil {
		return models.MapImageURL{}, nil
	}

	model := models.MapImageURL{
		Key:       fmt.Sprintf("%s-%d", m.Building, m.Number),
		MapDetail: m,
		URL:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.mapURLCollectionRepo.InsertOne(ctx, model); err != nil {
		return models.MapImageURL{}, nil
	}

	return model, nil
}

func (s *service) GetMapURLFromKey(ctx context.Context, floor int, building string, role constants.UserRole) (models.MapImageURL, error) {
	checkExistFilter := mongodb.Filter{"key": fmt.Sprintf("%s-%d", building, floor)}
	mapURL, err := s.mapURLCollectionRepo.FindOne(ctx, checkExistFilter)
	if err != nil {
		return models.MapImageURL{}, status.Error(codes.NotFound, "cannot found map url in database make sure your input is correct")
	}

	filter, _ := mongodb.AddFilter(mongodb.Filter{"building": building}, mongodb.Filter{"number": floor})
	mapDeatil, err := s.mapCollectionRepo.FindOne(ctx, filter)
	if err != nil {
		return models.MapImageURL{}, status.Error(codes.NotFound, "cannot found map detail in database make sure your input is correct")
	}

	if mapDeatil.IsAdmin && role != constants.AdminRole {
		return models.MapImageURL{}, status.Error(codes.PermissionDenied, "your not have permission to view this data")
	}

	return models.MapImageURL{
		Key:       mapURL.Key,
		MapDetail: mapDeatil,
		URL:       mapURL.URL,
		CreatedAt: mapDeatil.CreatedAt,
		UpdatedAt: mapDeatil.UpdatedAt,
	}, nil
}

func (s *service) AddFloorToDB(ctx context.Context, body models.Map) error {
	filter, _ := mongodb.AddFilter(mongodb.Filter{"building": body.Building}, mongodb.Filter{"number": body.Number})
	_, err := s.mapCollectionRepo.FindOne(ctx, filter)
	if err == nil {
		return status.Error(codes.AlreadyExists, "floor information is already exist in database")
	}

	if err := s.mapCollectionRepo.InsertOne(ctx, body); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *service) GetMapURLFromFlootNumber(ctx context.Context, floor int) (models.Map, error) {

	return models.Map{}, nil
}
