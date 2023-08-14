package signature

import (
	"net/url"
	"testing"
	"time"
)

const (
	key    = "sign"
	secret = "i1ydX9RtHyuJTrw7frcu"
	method = "POST"
	path   = "/echo"
	ttl    = time.Minute * 5
)

func TestSignature(t *testing.T) {
	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1 c2")
	params.Add("method", method)
	params.Add("path", path)

	sign := New(key, secret, ttl)
	authorization, date, err := sign.Generate(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("authorization:", authorization)
	t.Log("date:", date)
	ok, err := sign.Verify(authorization, date, params)
	t.Log(ok)
}

func TestSignatureSha256(t *testing.T) {
	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1 c2")
	params.Add("method", method)
	params.Add("path", path)

	sign := NewSha256(key, secret, ttl)
	authorization, date, err := sign.Generate(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("authorization:", authorization)
	t.Log("date:", date)
	ok, err := sign.Verify(authorization, date, params)
	t.Log(ok)
}
