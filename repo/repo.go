package repo

import (
	"context"
	"reflect"
	"time"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var g singleflight.Group

// Repo struct
type Repo struct {
	DB    *gorm.DB
	Cache cache.Cache
}

// GetDB 获取 gorm.DB
func (r *Repo) GetDB() *gorm.DB {
	return r.DB
}

// GetCache 获取 cache
func (r *Repo) GetCache() cache.Cache {
	return r.Cache
}

// Close mysql connection
func (r *Repo) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return nil
	}
	return db.Close()
}

// QueryCache 查询启用缓存
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) QueryCache(ctx context.Context, key string, data any, ttl time.Duration, query func(any) error) (err error) {
	// 从cache获取
	err = r.Cache.Get(ctx, key, data)
	if errors.Is(err, cache.ErrPlaceholder) {
		// 空数据也需要返回空的数据结构，保持与gorm返回一直的结构 see gorm.first()
		reflectValue := reflect.ValueOf(data)
		for reflectValue.Kind() == reflect.Ptr {
			if reflectValue.IsNil() && reflectValue.CanAddr() {
				reflectValue.Set(reflect.New(reflectValue.Type().Elem()))
			}

			reflectValue = reflectValue.Elem()
		}
		switch reflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			if reflectValue.Len() == 0 && reflectValue.Cap() == 0 {
				// if the slice cap is externally initialized, the externally initialized slice is directly used here
				reflectValue.Set(reflect.MakeSlice(reflectValue.Type(), 0, 20))
			}
		}
		logger.Debugf("[repo] key %v is empty", key)
		return nil
	} else if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "[repo] get cache by key: %s", key)
	}

	// 检查缓存取出的数据是否为空，不为空说明已经从缓存中取到了数据，直接返回
	if elem := reflect.ValueOf(data).Elem(); !elem.IsNil() {
		logger.Debugf("[repo] get from obj cache, key: %v, kind:%v", key, elem.Kind())
		return
	}

	// use sync/singleflight mode to get data
	// why not use redis lock? see this topic: https://redis.io/topics/distlock
	// demo see: https://github.com/go-demo/singleflight-demo/blob/master/main.go
	// https://juejin.cn/post/6844904084445593613
	_, err, _ = g.Do(key, func() (any, error) {
		// 从数据库中获取
		err = query(data)
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			if err = r.Cache.SetCacheWithNotFound(ctx, key); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound err, key: %s", key)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo] query db")
		}

		// set cache
		if err = r.Cache.Set(ctx, key, data, ttl); err != nil {
			return nil, errors.Wrapf(err, "[repo] set data to cache key: %s", key)
		}
		return data, nil
	})
	if err != nil {
		return errors.Wrapf(err, "[repo] get err via single flight do key: %s", key)
	}

	return nil
}

// DelCache 删除缓存
func (r *Repo) DelCache(ctx context.Context, key string) {
	if err := r.Cache.Del(ctx, key); err != nil {
		logger.Warnf("[repo] del cache key: %v", key)
	}
}
