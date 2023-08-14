package signature

import (
	"encoding/base64"
	"fmt"
	"time"
)

// Generate 生成签名
func (s *signature) Generate(params any) (auth string, ts int64, err error) {
	ts = time.Now().Unix()

	buffer, err := s.data(ts, params)
	if err != nil {
		return "", 0, err
	}
	digest := base64.StdEncoding.EncodeToString(s.cryptoFunc(buffer, []byte(s.secret)))

	auth = fmt.Sprintf("%s %s", s.key, digest)
	return auth, ts, nil
}
