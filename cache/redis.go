package cache

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/binbinly/pkg/codec"
	"github.com/binbinly/pkg/logger"
)

type redisCache struct {
	client *redis.Client
	codec  codec.Encoding
	expire time.Duration
	prefix string
}

// NewRedisCache new redis cache
func NewRedisCache(client *redis.Client, opts ...Option) Cache {
	o := NewOptions(opts...)
	return &redisCache{
		client: client,
		prefix: o.prefix,
		codec:  o.codec,
		expire: o.expire,
	}
}

// Set cache
func (c *redisCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	buf, err := c.codec.Marshal(val)
	if err != nil {
		return errors.Wrapf(err, "[cache] marshal data err, value is %+v", val)
	}

	if expiration == 0 {
		expiration = c.expire
	}

	if err = c.client.Set(ctx, c.buildKey(key), buf, expiration).Err(); err != nil {
		return errors.Wrapf(err, "[cache] redis set error")
	}
	return nil
}

// Get cache
func (c *redisCache) Get(ctx context.Context, key string, val any) error {
	cacheKey := c.buildKey(key)
	data, err := c.client.Get(ctx, cacheKey).Bytes()
	if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "[cache] get data error from redis, key is %+v", cacheKey)
	}

	// 防止data为空时，Unmarshal报错
	if string(data) == "" {
		return nil
	}
	if string(data) == NotFoundPlaceholder {
		return ErrPlaceholder
	}

	if err = c.codec.Unmarshal(data, val); err != nil {
		return errors.Wrapf(err, "[cache] unmarshal data error, key=%s, cacheKey=%s type=%v, data=%+v ",
			key, cacheKey, reflect.TypeOf(val), string(data))
	}
	return nil
}

// MultiSet 批量设置缓存
func (c *redisCache) MultiSet(ctx context.Context, m map[string]any, expiration time.Duration) error {
	if len(m) == 0 {
		return nil
	}
	if expiration == 0 {
		expiration = c.expire
	}
	// key-value是成对的，所以这里的容量是map的2倍
	paris := make([]any, 0, 2*len(m))
	for key, value := range m {
		buf, err := c.codec.Marshal(value)
		if err != nil {
			continue
		}
		paris = append(paris, []byte(c.buildKey(key)))
		paris = append(paris, buf)
	}

	if err := c.client.MSet(ctx, paris...).Err(); err != nil {
		return errors.Wrapf(err, "[cache] redis multi set error")
	}
	// 设置过期时间
	pipe := c.client.Pipeline()
	for i := 0; i < len(paris); i = i + 2 {
		switch paris[i].(type) {
		case []byte:
			pipe.Expire(ctx, string(paris[i].([]byte)), expiration)
		}
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return errors.Wrapf(err, "[cache] redis multi set expire error")
	}

	return nil
}

// MultiGet 批量获取缓存
func (c *redisCache) MultiGet(ctx context.Context, keys []string, value any, obj func() any) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKeys[index] = c.buildKey(key)
	}
	values, err := c.client.MGet(ctx, cacheKeys...).Result()
	if err != nil {
		return errors.Wrapf(err, "[cache] redis MGet error, keys is %+v", keys)
	}

	// 通过反射注入到map
	valueMap := reflect.ValueOf(value)
	for i, val := range values {
		if val == nil {
			continue
		}
		object := obj()
		if val.(string) == NotFoundPlaceholder {
			valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(object))
			continue
		}

		if err = c.codec.Unmarshal([]byte(val.(string)), &object); err != nil {
			logger.Warnf("[cache] unmarshal data error: %+v, key=%s, type=%v val=%v", err,
				keys[i], reflect.TypeOf(val), val)
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(object))
	}
	return nil
}

// Del cache
func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	// 批量构建cacheKey
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKeys[index] = c.buildKey(key)
	}

	if err := c.client.Del(ctx, cacheKeys...).Err(); err != nil {
		return errors.Wrapf(err, "[cache] redis delete error, keys is %+v", keys)
	}
	return nil
}

// SetCacheWithNotFound 设置空值
func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	return c.client.Set(ctx, c.buildKey(key), NotFoundPlaceholder, DefaultNotFoundExpireTime).Err()
}

func (c *redisCache) buildKey(key string) string {
	return strings.Join([]string{c.prefix, key}, ":")
}
