package handler

import (
	"context"

	"github.com/RyuChk/ips-map-service/cmd/map-grpc/internal/mapper"
	"github.com/RyuChk/ips-map-service/internal/constants"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	mapservice "github.com/RyuChk/ips-map-service/internal/services/mapService"
)

type MapV1Impl struct {
	mapService mapservice.Service
	mapv1.UnimplementedMapServiceServer
}

func ProvideMapServer(mapService mapservice.Service) mapv1.MapServiceServer {
	return &MapV1Impl{
		mapService: mapService,
	}
}

func (s *MapV1Impl) AddBuilding(ctx context.Context, req *mapv1.AddBuildingRequest) (*mapv1.AddBuildingResponse, error) {
	if err := s.mapService.AddBuildingInfoToDB(ctx, mapper.ToBuildingModel(req)); err != nil {
		return nil, err
	}

	return &mapv1.AddBuildingResponse{
		Name:        req.Name,
		Description: req.Description,
		OriginLat:   req.OriginLat,
		OriginLong:  req.OriginLong,
	}, nil
}

func (s *MapV1Impl) AddFloor(ctx context.Context, req *mapv1.AddFloorRequest) (*mapv1.AddFloorResponse, error) {
	if err := s.mapService.AddFloorToDB(ctx, mapper.ToFloorModel(req)); err != nil {
		return nil, err
	}

	return &mapv1.AddFloorResponse{}, nil
}

func (s *MapV1Impl) AddFloorDetail(ctx context.Context, req *mapv1.AddFloorDetailRequest) (*mapv1.AddFloorDetailResponse, error) {
	if err := s.mapService.UpsertFloorDetail(ctx, mapper.ToFloorDetailModel(req)); err != nil {
		return nil, err
	}

	return &mapv1.AddFloorDetailResponse{
		Building: req.Building,
		Floor:    req.Floor,
		Rooms:    req.Rooms,
	}, nil
}

func (s *MapV1Impl) GetBuildingList(ctx context.Context, req *mapv1.GetBuildingListRequest) (*mapv1.GetBuildingListResponse, error) {
	list, err := s.mapService.GetBuildingList(ctx, mapper.ToRole[req.Role] == constants.AdminRole)
	if err != nil {
		return nil, err
	}

	return mapper.ToGetBuildingListResponse(list), nil
}

func (s *MapV1Impl) GetBuildingInfo(ctx context.Context, req *mapv1.GetBuildingInfoRequest) (*mapv1.GetBuildingInfoResponse, error) {
	resp, err := s.mapService.GetBuildingInfo(ctx, req.Building, mapper.ToRole[req.Role])
	if err != nil {
		return nil, err
	}

	return mapper.ToGetBuildingInfoResponse(resp), nil
}

func (s *MapV1Impl) GetFloorInfo(ctx context.Context, req *mapv1.GetFloorInfoRequest) (*mapv1.GetFloorInfoResponse, error) {
	resp, err := s.mapService.GetFloorDetailFromKey(ctx, int(req.FloorNumber), req.Building, mapper.ToRole[req.Role])
	if err != nil {
		return nil, err
	}

	return mapper.ToGetFloorInfoResponse(resp), nil
}
