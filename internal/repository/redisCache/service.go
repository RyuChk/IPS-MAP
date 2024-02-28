package rediscache

import (
	"context"
	"time"

	wireredis "git.cie-ips.com/ips/wire-provider/redis"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Service interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Scan(ctx context.Context, prefix string) ([][]byte, error)
}

type service struct {
	client *redis.Client
}

func ProvideCacheService(cfg wireredis.Config) (Service, error) {
	conn, err := wireredis.NewConnection(&cfg)
	if err != nil {
		return nil, nil
	}
	if conn == nil {
		panic(err)
	}
	return &service{
		client: conn.Client(),
	}, nil
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	res, err := s.client.Get(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).Str("key", key).Msg("error getting redis key")
		return "", err
	}

	return res, nil
}

func (s *service) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := s.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		log.Error().Err(err).Str("key", key).Any("value", value).Dur("ttl", ttl).Msg("error setting redis key")
		return err
	}

	return nil
}

func (s *service) Scan(ctx context.Context, prefix string) ([][]byte, error) {
	iter := s.client.Scan(ctx, 0, prefix, 0).Iterator()
	if iter.Err() != nil {
		log.Error().Err(iter.Err()).Str("prefix", prefix).Msg("error scan")
		return nil, iter.Err()
	}

	result := make([][]byte, 0)
	for iter.Next(ctx) {
		val, err := s.Get(ctx, iter.Val())
		if err != nil {
			log.Error().Err(err).Str("key", iter.Val()).Msg("error getting redis key")
			continue
		}
		result = append(result, []byte(val))
	}

	return result, nil
}
