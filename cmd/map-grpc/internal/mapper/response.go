package mapper

import (
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	"github.com/RyuChk/ips-map-service/internal/models"
)

func ToGetBuildingListResponse(body []models.Building) *mapv1.GetBuildingListResponse {
	result := &mapv1.GetBuildingListResponse{
		Buildings: make([]*mapv1.Building, 0),
	}

	for _, v := range body {
		result.Buildings = append(result.Buildings, &mapv1.Building{
			Name:        v.Name,
			Description: v.Description,
			OriginLat:   v.OriginLat,
			OriginLong:  v.OriginLong,
			IsAdmin:     false,
		})
	}

	return result
}

func ToGetFloorInfoResponse(body models.FloorDetail) *mapv1.GetFloorInfoResponse {
	result := &mapv1.GetFloorInfoResponse{
		Info: &mapv1.FloorDetail{
			Name:        body.Info.Name,
			Description: body.Info.Description,
			Floor:       int32(body.Info.Floor),
			Building:    body.Info.Building,
			Symbol:      body.Info.Symbol,
			IsAdmin:     body.Info.IsAdmin,
			OriginLat:   body.Info.OriginLat,
			OriginLong:  body.Info.OriginLong,
		},
		Rooms: make([]*mapv1.RoomDetail, 0),
	}

	for _, v := range body.RoomList {
		result.Rooms = append(result.Rooms, &mapv1.RoomDetail{
			RoomId:      v.RoomID,
			Name:        v.Name,
			Description: v.Description,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
		})
	}

	return result
}

func ToGetBuildingInfoResponse(body models.Building) *mapv1.GetBuildingInfoResponse {
	result := &mapv1.GetBuildingInfoResponse{
		Name:        body.Name,
		Description: body.Description,
		OriginLat:   body.OriginLat,
		OriginLong:  body.OriginLong,
		Floors:      make([]*mapv1.FloorDetail, len(body.FloorList)),
		IsAdmin:     false,
	}

	for i, v := range body.FloorList {
		result.Floors[i] = &mapv1.FloorDetail{
			Name:        v.Name,
			Description: v.Description,
			Floor:       int32(v.Floor),
			Building:    v.Building,
			Symbol:      v.Symbol,
			IsAdmin:     v.IsAdmin,
			OriginLat:   v.OriginLat,
			OriginLong:  v.OriginLong,
		}
	}

	return result
}
