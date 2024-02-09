package config

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"git.cie.com/ips/wire-provider/redis"
	"github.com/kelseyhightower/envconfig"
)

func ProvideMongoxConfig() wiremongo.Config {
	return provideConfig(wiremongo.Config{})
}

func ProvideMinioXConfig() wireminio.Config {
	return provideConfig(wireminio.Config{})
}

func ProvideCacheConfig() CacheConfig {
	return provideConfig(CacheConfig{})
}

func ProvideMinioConfig() MinioConfig {
	return provideConfig(MinioConfig{})
}

func ProvideUserTrackingConfig() UserTrackingConfig {
	return provideConfig(UserTrackingConfig{})
}

func ProvideRedisCacheConfig() redis.Config {
	return provideConfig(redis.Config{})
}

func provideConfig[T any](cfg T) T {
	envconfig.MustProcess("", &cfg)
	return cfg
}
