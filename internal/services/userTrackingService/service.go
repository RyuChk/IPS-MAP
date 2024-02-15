package usertrackingservice

import (
	"context"
	"fmt"

	"github.com/RyuChk/ips-map-service/internal/config"
	"github.com/RyuChk/ips-map-service/internal/constants"
	"github.com/RyuChk/ips-map-service/internal/models"
	rediscache "github.com/RyuChk/ips-map-service/internal/repository/redisCache"
	"github.com/RyuChk/ips-map-service/utils"
	"github.com/rs/zerolog/log"
)

type Service interface {
	FetchOnlineUserFromFilter(ctx context.Context, building string, floor int) (*[]models.OnlineUser, error)
	AddUpdateOnlineUser(ctx context.Context, userInfo models.OnlineUser) error
}

type service struct {
	cfg         config.UserTrackingConfig
	redisClient rediscache.Service
}

func ProvideUserTrackingService(cfg config.UserTrackingConfig, redisClient rediscache.Service) Service {
	return &service{
		redisClient: redisClient,
		cfg:         cfg,
	}
}

func prefixGenerator(name, building string, floor int) string {
	return fmt.Sprintf("%s%s:%d:%s", constants.UserTrackingPrefix, building, floor, name)
}

func (s *service) FetchOnlineUserFromFilter(ctx context.Context, building string, floor int) (*[]models.OnlineUser, error) {
	prefix := fmt.Sprintf("%s%s:%d:*", constants.UserTrackingPrefix, building, floor)
	resp, err := s.redisClient.Scan(ctx, prefix)
	if err != nil {
		return nil, err
	}

	result := make([]models.OnlineUser, len(resp))
	for i, v := range resp {
		var temp models.OnlineUser
		if err := utils.FromByte(v, &temp); err != nil {
			return nil, err
		}
		result[i] = temp
	}

	return &result, err
}

func (s *service) AddUpdateOnlineUser(ctx context.Context, userInfo models.OnlineUser) error {
	bytes, err := utils.ToByte(userInfo)
	if err != nil {
		log.Error().Any("user_info", userInfo).Msg("error convert to byte array")
	}

	var value interface{} = bytes
	prefix := prefixGenerator(userInfo.DisplayName, userInfo.Building, userInfo.Floor)
	if err := s.redisClient.Set(ctx, prefix, value, s.cfg.UserTrackingTTL); err != nil {
		return err
	}
	return nil
}
