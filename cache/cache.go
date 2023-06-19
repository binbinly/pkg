package cache

import (
	"context"
	"errors"
	"time"
)

const (
	// DefaultPrefix 默认缓存前缀
	DefaultPrefix = "cache:"
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime 结果为空时的过期时间 1分钟, 常用于数据为空时的缓存时间(缓存穿透)
	DefaultNotFoundExpireTime = time.Minute
	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"
)

var (
	// ErrPlaceholder 空数据标识
	ErrPlaceholder = errors.New("cache: placeholder")
)

// Cache 定义cache驱动接口
type Cache interface {
	Set(ctx context.Context, key string, val any, expiration time.Duration) error
	Get(ctx context.Context, key string, val any) error
	MultiSet(ctx context.Context, valMap map[string]any, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap any, newObject func() any) error
	Del(ctx context.Context, keys ...string) error
	SetCacheWithNotFound(ctx context.Context, key string) error
}
