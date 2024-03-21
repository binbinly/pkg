package cache

import (
	"context"
	"errors"
	"strings"
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
	// ErrSetMemoryWithNotFound 设置内存缓存时key不存在
	ErrSetMemoryWithNotFound = errors.New("cache: set memory cache err for not found")
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

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(prefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey = key
	if prefix != "" {
		cacheKey, err = strings.Join([]string{prefix, key}, ":"), nil
	}

	return
}
