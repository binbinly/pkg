package lock

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	_prefix = "lock:"
	_ttl    = 30 * time.Second
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
