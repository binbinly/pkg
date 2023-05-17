package redis

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var Client *redis.Client

func TestMain(m *testing.M) {
	Client = InitTestRedis()
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestInitTestRedis(t *testing.T) {
	err := Client.Ping(context.Background()).Err()
	assert.Nil(t, err)
}

func TestRedisSetGet(t *testing.T) {
	var key = "test-set"
	var value = "test-content"
	Client.Set(context.Background(), key, value, time.Second*100)

	expectValue := Client.Get(context.Background(), key).Val()
	assert.Equal(t, expectValue, value)
}

func BenchmarkGoroutineData(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			Client.Set(context.Background(), fmt.Sprintf("test-%d", i), "test-content", time.Minute)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
}
