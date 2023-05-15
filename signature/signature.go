package signature

import (
	"net/url"
	"time"
)

var _ Signature = (*signature)(nil)

// CryptoFunc 签名加密函数
// type CryptoFunc func(b []byte, key []byte, h func() hash.Hash) []byte

const (
	delimiter = "|"
)

type Signature interface {
	// Generate 生成签名
	Generate(params url.Values) (authorization, date string, err error)
	// Verify 验证签名
	Verify(authorization, date string, params url.Values) (ok bool, err error)
}

type signature struct {
	key    string
	secret string
	ttl    time.Duration
	// cryptoFunc CryptoFunc
}

func New(key, secret string, ttl time.Duration) Signature {
	return &signature{
		key:    key,
		secret: secret,
		ttl:    ttl,
	}
}
