package handler

import (
	"context"

	"github.com/RyuChk/ips-map-service/cmd/map-grpc/internal/mapper"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	mapservice "github.com/RyuChk/ips-map-service/internal/services/mapService"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *MapV1Impl) AddFloor(ctx context.Context, req *mapv1.AddFloorRequest) (*mapv1.AddFloorResponse, error) {
	if err := s.mapService.AddFloorToDB(ctx, mapper.ToMapModel(req)); err != nil {
		return &mapv1.AddFloorResponse{}, err
	}
	return &mapv1.AddFloorResponse{
		Info:      (*mapv1.FloorDetail)(req),
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (h *MapV1Impl) GetFloorList(ctx context.Context, req *mapv1.GetFloorListRequest) (*mapv1.GetFloorListResponse, error) {
	floors, err := h.mapService.GetFloorListByBuilding(ctx, req.Building)
	if err != nil {
		return nil, err
	}

	return mapper.ToGetFloorListResponse(floors), nil
}

func (h *MapV1Impl) AddMapURL(ctx context.Context, req *mapv1.AddMapURLRequest) (*mapv1.AddMapURLResponse, error) {
	m, err := h.mapService.AddMapURL(ctx, int(req.FloorNumber), req.Building, req.Url)
	if err != nil {
		return nil, err
	}
	return mapper.ToAddMapURLResponse(m), nil
}

func (h *MapV1Impl) GetMapURL(ctx context.Context, req *mapv1.GetMapURLRequest) (*mapv1.GetMapURLResponse, error) {
	m, err := h.mapService.GetMapURLFromKey(ctx, int(req.FloorNumber), req.Building)
	if err != nil {
		return nil, err
	}

	return mapper.ToGetMapURLResponse(m), nil
}
