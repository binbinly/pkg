package lock

import (
	"context"
	"time"

	"github.com/google/uuid"
)

var (
	_defaultPrefix = "lock:"
	_defaultTTL    = 30 * time.Second
)

// Lock define common func
type Lock interface {
	Lock(ctx context.Context) (bool, error)
	Unlock(ctx context.Context) (bool, error)
}

// genToken 生成token
func genToken() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func buildKey(prefix, key string) string {
	return prefix + key
}
