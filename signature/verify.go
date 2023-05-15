package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/binbinly/pkg/util"
)

// Verify verifies the signature
func (s *signature) Verify(auth, date string, params url.Values) (bool, error) {
	if date == "" {
		return false, errors.New("date required")
	}

	timestamp, err := util.ParseCSTInLocation(date)
	if err != nil {
		return false, err
	}

	if util.SubInLocation(timestamp) > float64(s.ttl/time.Second) {
		return false, errors.New("date exceeds limit")
	}

	// The Encode() method already sorts by key
	sortedParamsEncoded, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return false, err
	}

	buffer := bytes.NewBufferString(sortedParamsEncoded)
	buffer.WriteString(delimiter)
	buffer.WriteString(date)

	// Encrypt data using hmac and base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return auth == fmt.Sprintf("%s %s", s.key, digest), nil
}
