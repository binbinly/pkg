package signature

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/binbinly/pkg/util"
)

// Verify verifies the signature
func (s *signature) Verify(auth string, timestamp int64, params any) (bool, error) {
	if timestamp == 0 {
		return false, errors.New("date required")
	}
	ts := time.Unix(timestamp, 0)
	if util.SubInLocation(ts) > float64(s.ttl/time.Second) {
		return false, errors.New("date exceeds limit")
	}

	buffer, err := s.data(timestamp, params)
	if err != nil {
		return false, err
	}

	digest := base64.StdEncoding.EncodeToString(s.cryptoFunc(buffer, []byte(s.secret)))

	return auth == fmt.Sprintf("%s %s", s.key, digest), nil
}
