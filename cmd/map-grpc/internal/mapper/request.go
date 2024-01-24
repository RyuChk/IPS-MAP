package mapper

import (
	"time"

	"github.com/RyuChk/ips-map-service/apps/map/models"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
)

func ToMapModel(req *mapv1.AddFloorRequest) models.Map {
	return models.Map{
		Name:        req.Name,
		Description: req.Description,
		Symbol:      req.Symbol,
		Building:    req.Building,
		Number:      int(req.Number),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
