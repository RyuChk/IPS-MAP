package handler

import (
	"context"

	"github.com/RyuChk/ips-map-service/cmd/user-tracking-grpc/internal/mapper"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	usertrackingservice "github.com/RyuChk/ips-map-service/internal/services/userTrackingService"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserTrackingV1Impl struct {
	userTrackingService usertrackingservice.Service
	mapv1.UnimplementedUserTrackingServiceServer
}

func ProvideUserTrackingServer(userTrackingService usertrackingservice.Service) mapv1.UserTrackingServiceServer {
	return &UserTrackingV1Impl{
		userTrackingService: userTrackingService,
	}
}

func (h *UserTrackingV1Impl) FetchOnlineUser(ctx context.Context, req *mapv1.FetchOnlineUserRequest) (*mapv1.FetchOnlineUserResponse, error) {
	onlineUser, err := h.userTrackingService.FetchOnlineUserFromFilter(ctx, req.Building, int(req.Floor))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return mapper.MapOnlineUserToFetchOnlineResponse(onlineUser), nil
}

func (h *UserTrackingV1Impl) AddUpdateOnlineUser(ctx context.Context, req *mapv1.AddUpdateOnlineUserRequest) (*mapv1.AddUpdateOnlineUserResponse, error) {
	if err := h.userTrackingService.AddUpdateOnlineUser(ctx, mapper.MapRequestToOnlineUserModel(req)); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &mapv1.AddUpdateOnlineUserResponse{}, nil
}
