package constants

type CollectionStage string

type RedisCachePrefix string

var (
	MapURLCachePrefix RedisCachePrefix = "MAP:IMAGE:FLOOR:"
)

type UserRole string

var (
	AdminRole UserRole = "ADMIN"
	User      UserRole = "USER"
)
