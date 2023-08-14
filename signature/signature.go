package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"net/url"
	"strconv"
	"time"

	"github.com/binbinly/pkg/util"
	"github.com/binbinly/pkg/util/xhash"
)

var _ Signature = (*signature)(nil)

// CryptoFunc 签名加密函数
type CryptoFunc func(b, k []byte) []byte

const (
	delimiter = "|"
)

type Signature interface {
	// Generate 生成签名
	Generate(params any) (auth string, ts int64, err error)
	// Verify 验证签名
	Verify(auth string, ts int64, params any) (ok bool, err error)
}

type signature struct {
	key        string
	secret     string
	ttl        time.Duration
	cryptoFunc CryptoFunc
}

func New(key, secret string, ttl time.Duration) Signature {
	return &signature{
		key:    key,
		secret: secret,
		ttl:    ttl,
		cryptoFunc: func(b, k []byte) []byte {
			buf := bytes.NewBuffer(b)
			buf.WriteString("&key=")
			buf.Write(k)
			return xhash.MD5(buf.Bytes())
		},
	}
}

func NewSha256(key, secret string, ttl time.Duration) Signature {
	return &signature{
		key:    key,
		secret: secret,
		ttl:    ttl,
		cryptoFunc: func(b, k []byte) []byte {
			hash := hmac.New(sha256.New, k)
			_, _ = hash.Write(b)
			return hash.Sum(nil)
		},
	}
}

func NewCrypto(key, secret string, ttl time.Duration, f CryptoFunc) Signature {
	return &signature{
		key:        key,
		secret:     secret,
		ttl:        ttl,
		cryptoFunc: f,
	}
}

func (s *signature) data(timestamp int64, params any) ([]byte, error) {
	buffer := bytes.NewBufferString(strconv.FormatInt(timestamp, 10))
	buffer.WriteString(delimiter)
	switch p := params.(type) {
	case url.Values:
		// Encode() 方法中自带 sorted by key
		sortParamsEncode, err := url.QueryUnescape(p.Encode())
		if err != nil {
			return nil, err
		}
		buffer.WriteString(sortParamsEncode)
	case map[string]any:
		buffer.WriteString(util.MapBuildQuery(p))
	case string:
		buffer.WriteString(p)
	}

	return buffer.Bytes(), nil
}
