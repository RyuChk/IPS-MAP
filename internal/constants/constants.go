package constants

type CollectionStage string

type RedisCachePrefix string

var (
	MapURLCachePrefix  RedisCachePrefix = "MAP:IMAGE:FLOOR:"
	UserTrackingPrefix RedisCachePrefix = "TRACKING:"
)

func (r RedisCachePrefix) String() string {
	return string(r)
}

type UserRole string

var (
	AdminRole UserRole = "ADMIN"
	User      UserRole = "USER"
)
