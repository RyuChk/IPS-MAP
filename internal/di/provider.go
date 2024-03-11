package di

import (
	wireminio "git.cie-ips.com/ips/wire-provider/minio"
	wiremongo "git.cie-ips.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/internal/config"
	"github.com/RyuChk/ips-map-service/internal/repository/cache"
	"github.com/RyuChk/ips-map-service/internal/repository/minio"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	buildingcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/buildingCollectionRepo"
	floorcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/floorCollectionRepo"
	floordetailcollectionrepo "github.com/RyuChk/ips-map-service/internal/repository/mongodb/floorDetailCollectionRepo"
	rediscache "github.com/RyuChk/ips-map-service/internal/repository/redisCache"
	mapservice "github.com/RyuChk/ips-map-service/internal/services/mapService"
	usertrackingservice "github.com/RyuChk/ips-map-service/internal/services/userTrackingService"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
	rediscache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	buildingcollectionrepo.ProvideBuildingCollectionRepo,
	floorcollectionrepo.ProvideFloorCollectionRepo,
	floordetailcollectionrepo.ProvideMapURLCollectionRepo,
	mapservice.ProvideMapURLService,
	usertrackingservice.ProvideUserTrackingService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideMinioConfig,
	config.ProvideUserTrackingConfig,
	config.ProvideRedisCacheConfig,
)

type Locator struct {
	MongoDBConn  wiremongo.Connection
	MinioXConn   wireminio.Connection
	CacheService cache.Service
	MapService   mapservice.Service
}
