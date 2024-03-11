package mapper

import (
	"fmt"
	"time"

	"github.com/RyuChk/ips-map-service/internal/constants"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	userv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/user/v1"
	"github.com/RyuChk/ips-map-service/internal/models"
)

var ToRole = map[userv1.Role]constants.UserRole{
	userv1.Role_ROLE_ADMIN: constants.AdminRole,
	userv1.Role_ROLE_USER:  constants.User,
}

func ToFloorModel(req *mapv1.AddFloorRequest) models.Floor {
	return models.Floor{
		Name:        req.Name,
		Description: req.Description,
		Symbol:      req.Symbol,
		Building:    req.Building,
		IsAdmin:     req.IsAdmin,
		Floor:       int(req.Number),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToBuildingModel(req *mapv1.AddBuildingRequest) models.Building {
	return models.Building{
		Name:        req.Name,
		Description: req.Description,
		OriginLat:   req.OriginLat,
		OriginLong:  req.OriginLong,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToFloorDetailModel(req *mapv1.AddFloorDetailRequest) models.FloorDetail {
	result := models.FloorDetail{
		Key:       fmt.Sprintf("%s-%d", req.Building, req.Floor),
		RoomList:  make([]models.Room, len(req.Rooms)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	for i, v := range req.Rooms {
		result.RoomList[i] = models.Room{
			RoomID:      v.RoomId,
			Name:        v.Name,
			Description: v.Description,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}

	return result
}
