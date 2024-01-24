package mapper

import (
	"github.com/RyuChk/ips-map-service/apps/map/models"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToGetFloorListResponse(body []models.Map) *mapv1.GetFloorListResponse {
	result := &mapv1.GetFloorListResponse{
		Floors: make([]*mapv1.FloorDetail, len(body)),
	}

	for i, v := range body {
		result.Floors[i] = &mapv1.FloorDetail{
			Name:        v.Name,
			Description: v.Description,
			Building:    v.Building,
			Symbol:      v.Symbol,
			Number:      int32(v.Number),
		}
	}

	return result
}

func ToAddMapURLResponse(body models.MapImageURL) *mapv1.AddMapURLResponse {
	return &mapv1.AddMapURLResponse{
		Detail: &mapv1.FloorDetail{
			Name:        body.MapDetail.Name,
			Description: body.MapDetail.Description,
			Building:    body.MapDetail.Building,
			Number:      int32(body.MapDetail.Number),
			Symbol:      body.MapDetail.Symbol,
		},
		Url:       body.URL,
		UpdatedAt: timestamppb.New(body.UpdatedAt),
		CreatedAt: timestamppb.New(body.CreatedAt),
	}
}

func ToGetMapURLResponse(body models.MapImageURL) *mapv1.GetMapURLResponse {
	return &mapv1.GetMapURLResponse{
		Detail: &mapv1.FloorDetail{
			Name:        body.MapDetail.Name,
			Description: body.MapDetail.Description,
			Building:    body.MapDetail.Building,
			Number:      int32(body.MapDetail.Number),
			Symbol:      body.MapDetail.Symbol,
		},
		Url:       body.URL,
		UpdatedAt: timestamppb.New(body.UpdatedAt),
	}
}
