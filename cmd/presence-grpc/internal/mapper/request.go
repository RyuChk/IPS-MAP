package mapper

import (
	"time"

	"github.com/RyuChk/ips-map-service/apps/constants"
	"github.com/RyuChk/ips-map-service/apps/map/models"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	userv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/user/v1"
)

var ToRole = map[userv1.Role]constants.UserRole{
	userv1.Role_ROLE_ADMIN: constants.AdminRole,
	userv1.Role_ROLE_USER:  constants.User,
}

func ToMapModel(req *mapv1.AddFloorRequest) models.Map {
	return models.Map{
		Name:        req.Name,
		Description: req.Description,
		Symbol:      req.Symbol,
		Building:    req.Building,
		IsAdmin:     req.IsAdmin,
		Number:      int(req.Number),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
