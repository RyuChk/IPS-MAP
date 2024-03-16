package mapper

import (
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	sharedmapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/map/v1"
	sharedrssiv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/rssi/v1"
	"github.com/RyuChk/ips-map-service/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapOnlineUserToFetchOnlineResponse(users *[]models.OnlineUser) *mapv1.FetchOnlineUserResponse {
	onlineUser := make([]*sharedmapv1.OnlineUser, len(*users))
	for i, v := range *users {
		onlineUser[i] = &sharedmapv1.OnlineUser{
			DisplayName: v.DisplayName,
			Coordinate: &sharedrssiv1.Position{
				X: float32(v.Coordinate.X),
				Y: float32(v.Coordinate.Y),
				Z: float32(v.Coordinate.Z),
			},
			Label:     v.Label,
			Building:  v.Building,
			Floor:     int32(v.Floor),
			Timestamp: timestamppb.New(v.Timestamp),
		}
	}

	return &mapv1.FetchOnlineUserResponse{
		OnlineUsers: onlineUser,
		Timestamp:   timestamppb.Now(),
	}
}
