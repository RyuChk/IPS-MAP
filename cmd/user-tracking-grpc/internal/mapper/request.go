package mapper

import (
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	"github.com/RyuChk/ips-map-service/internal/models"
)

func MapRequestToOnlineUserModel(req *mapv1.AddUpdateOnlineUserRequest) models.OnlineUser {
	return models.OnlineUser{
		DisplayName: req.DisplayName,
		Building:    req.Building,
		Floor:       int(req.Floor),
		Coordinate: models.Position{
			X: float64(req.Position.X),
			Y: float64(req.Position.Y),
			Z: float64(req.Position.Z),
		},
		Timestamp: req.Timestamp.AsTime(),
	}
}
