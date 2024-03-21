package cache

import (
	"context"
	"reflect"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

type memoryCache struct {
	client *ristretto.Cache
	opts   Options
}

// NewMemoryCache create a memory cache
func NewMemoryCache(opts ...Option) Cache {
	o := NewOptions(opts...)
	// see: https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/
	//		https://www.start.io/blog/we-chose-ristretto-cache-for-go-heres-why/
	config := &ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	}
	store, _ := ristretto.NewCache(config)
	return &memoryCache{
		client: store,
		opts:   o,
	}
}

// Set add cache
func (m *memoryCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	buf, err := m.opts.codec.Marshal(val)
	if err != nil {
		return errors.Wrapf(err, "[cache.menery] marshal data err, value is %+v", val)
	}
	cacheKey, err := BuildCacheKey(m.opts.prefix, key)
	if err != nil {
		return err
	}
	m.client.SetWithTTL(cacheKey, buf, 0, expiration)
	return nil
}

// Get data
func (m *memoryCache) Get(ctx context.Context, key string, val any) error {
	cacheKey, err := BuildCacheKey(m.opts.prefix, key)
	if err != nil {
		return err
	}
	data, ok := m.client.Get(cacheKey)
	if !ok {
		return nil
	}
	if data == NotFoundPlaceholder {
		return ErrPlaceholder
	}

	if err = m.opts.codec.Unmarshal(data.([]byte), val); err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data.([]byte)))
	}
	return nil
}

// Del 删除
func (m *memoryCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	cacheKey, err := BuildCacheKey(m.opts.prefix, key)
	if err != nil {
		return err
	}
	m.client.Del(cacheKey)
	return nil
}

// MultiSet 批量set
func (m *memoryCache) MultiSet(ctx context.Context, valMap map[string]any, expiration time.Duration) error {
	panic("implement me")
}

// MultiGet 批量获取
func (m *memoryCache) MultiGet(ctx context.Context, keys []string, valueMap any, newObject func() any) error {
	panic("implement me")
}

func (m *memoryCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	if m.client.Set(key, NotFoundPlaceholder, int64(DefaultNotFoundExpireTime)) {
		return nil
	}
	return ErrSetMemoryWithNotFound
}
