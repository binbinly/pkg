package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/binbinly/pkg/util"
)

// Generate 生成签名
func (s *signature) Generate(params url.Values) (authorization, date string, err error) {
	date = util.CSTLayoutString()

	// Encode() 方法中自带 sorted by key
	sortParamsEncode, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", "", err
	}

	// 加密字符串规则
	buffer := bytes.NewBufferString(sortParamsEncode)
	buffer.WriteString(delimiter)
	buffer.WriteString(date)

	// 对数据进行 sha256 加密，并进行 base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	if _, err := hash.Write(buffer.Bytes()); err != nil {
		return "", "", err
	}
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	authorization = fmt.Sprintf("%s %s", s.key, digest)
	return authorization, date, nil
}
