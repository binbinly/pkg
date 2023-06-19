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

// NewBasicClient new a redis instance
func NewBasicClient(addr, pwd string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})

	// check redis if is ok
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	log.Println("init redis success by addr:", addr)
	return rdb, nil
}

// NewClient new a redis instance
func NewClient(c *Config) (*redis.Client, error) {
	// create a redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})

	// check redis if is ok
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	// hook tracing (using open telemetry)
	if c.Trace {
		if err := redisotel.InstrumentTracing(rdb); err != nil {
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
