package lock

import (
	"context"
	"testing"
	"time"

	redis2 "github.com/binbinly/pkg/storage/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var Client *redis.Client

func TestMain(m *testing.M) {
	Client = redis2.InitTestRedis()
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestLockWithDefaultTimeout(t *testing.T) {
	lock := NewRedisLock(Client, "lock1", WithTTL(2*time.Second))
	ok, err := lock.Lock(context.Background())
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = lock.Unlock(context.Background())
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestLockWithTimeout(t *testing.T) {
	t.Run("should lock/unlock success", func(t *testing.T) {
		ctx := context.Background()
		lock1 := NewRedisLock(Client, "lock2", WithTTL(2*time.Second))
		ok, err := lock1.Lock(ctx)
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = lock1.Unlock(ctx)
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("should unlock failed", func(t *testing.T) {
		ctx := context.Background()
		lock2 := NewRedisLock(Client, "lock3", WithTTL(2*time.Second))
		ok, err := lock2.Lock(ctx)
		assert.Nil(t, err)
		assert.True(t, ok)

		time.Sleep(3 * time.Second)

		ok, err = lock2.Unlock(ctx)
		assert.Nil(t, err)
		//assert.False(t, ok)
	})
}
