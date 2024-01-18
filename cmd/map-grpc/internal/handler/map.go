package handler

import (
	"context"

	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
)

type MapV1Impl struct {
	mapv1.UnimplementedMapServiceServer
}

func ProvideMapServer() mapv1.MapServiceServer {
	return &MapV1Impl{}
}

func (s *MapV1Impl) GetMapURL(ctx context.Context, req *mapv1.GetMapURLRequest) (*mapv1.GetMapURLResponse, error) {
	return &mapv1.GetMapURLResponse{}, nil
}
