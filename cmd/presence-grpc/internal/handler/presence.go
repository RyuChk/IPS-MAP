package handler

import (
	"context"

	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	sharedv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/map/v1"
	coorsharedv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/rssi/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PresenceV1Impl struct {
	mapv1.UnimplementedPresenceServiceServer
}

func ProvidePresenceServer() mapv1.PresenceServiceServer {
	return &PresenceV1Impl{}
}

func (h *PresenceV1Impl) FetchOnlineUser(ctx context.Context, req *mapv1.FetchOnlineUserRequest) (*mapv1.FetchOnlineUserResponse, error) {
	return &mapv1.FetchOnlineUserResponse{
		OnlineUsers: []*sharedv1.OnlineUser{
			{
				DisplayName: "Mock User 1",
				Coordinate: &coorsharedv1.Position{
					X: 1,
					Y: 2,
					Z: 0,
				},
				Timestamp: timestamppb.Now(),
			},
			{
				DisplayName: "Mock User 2",
				Coordinate: &coorsharedv1.Position{
					X: 3,
					Y: 2,
					Z: 0,
				},
				Timestamp: timestamppb.Now(),
			},
		},
		Timestamp: timestamppb.Now(),
	}, nil
}
