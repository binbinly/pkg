package redis

import (
	"context"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

const (
	// Nil redis nil
	Nil = redis.Nil
	// Success redis成功标识
	Success = 1
)

// NewClient new a redis instance
func NewClient(c *Config) (rdb *redis.Client, err error) {
	if c.Url != "" {
		opt, err := redis.ParseURL(c.Url)
		if err != nil {
			return nil, err
		}
		rdb = redis.NewClient(opt)
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:         c.Addr,
			Username:     c.Username,
			Password:     c.Password,
			DB:           c.DB,
			MinIdleConns: c.MinIdleConn,
			DialTimeout:  c.DialTimeout,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			PoolSize:     c.PoolSize,
			PoolTimeout:  c.PoolTimeout,
		})
	}
	// check redis if is ok
	if _, err = rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	// hook tracing (using open telemetry)
	if c.Trace {
		if err = redisotel.InstrumentTracing(rdb); err != nil {
			return nil, err
		}
	}

	log.Println("init redis success by addr:", c.Addr)
	return rdb, nil
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	log.Println("mini redis addr:", mr.Addr())
	return rdb
}
