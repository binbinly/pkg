package lock

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// RedisLock is a redis lock.
type RedisLock struct {
	prefix string
	key    string
	token  string
	rdb    *redis.Client
	ttl    time.Duration
}

// Option RedisLock option
type Option func(l *RedisLock)

// WithPrefix with prefix
func WithPrefix(prefix string) Option {
	return func(l *RedisLock) {
		l.prefix = prefix
	}
}

// WithTTL with ttl
func WithTTL(ttl time.Duration) Option {
	return func(l *RedisLock) {
		l.ttl = ttl
	}
}

// NewRedisLock new a redis lock instance
func NewRedisLock(rdb *redis.Client, key string, opts ...Option) *RedisLock {
	opt := &RedisLock{
		rdb:    rdb,
		token:  genToken(),
		prefix: _defaultPrefix,
		ttl:    _defaultTTL,
	}
	for _, f := range opts {
		f(opt)
	}
	opt.key = buildKey(opt.prefix, key)
	return opt
}

// Lock acquires the lock.
func (l *RedisLock) Lock(ctx context.Context) (bool, error) {
	isSet, err := l.rdb.SetNX(ctx, l.key, l.token, l.ttl).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, errors.Wrapf(err, "[lock] acquires the lock err, key: %s", l.key)
	}
	return isSet, nil
}

// Unlock del the lock.
// NOTE: token 一致才会执行删除，避免误删，这里用了lua脚本进行事务处理
func (l *RedisLock) Unlock(ctx context.Context) (bool, error) {
	luaScript := "if redis.call('GET',KEYS[1]) == ARGV[1] then return redis.call('DEL',KEYS[1]) else return 0 end"
	ret, err := l.rdb.Eval(ctx, luaScript, []string{l.key}, l.token).Result()
	if err != nil {
		return false, err
	}
	reply, ok := ret.(int64)
	if !ok {
		return false, nil
	}
	return reply == 1, nil
}
