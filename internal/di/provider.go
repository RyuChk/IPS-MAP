package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/internal/config"
	"github.com/RyuChk/ips-map-service/internal/repository/cache"
	"github.com/RyuChk/ips-map-service/internal/repository/minio"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	mapurlcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/mapURLCollectionRepo"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	mapurlcollectionrepo.ProvideMapURLCollectionRepo,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
)

type Locator struct {
	MongoDBConn  wiremongo.Connection
	MinioXConn   wireminio.Connection
	CacheService cache.Service
}
