package config

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
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

func provideConfig[T any](cfg T) T {
	envconfig.MustProcess("", &cfg)
	return cfg
}
