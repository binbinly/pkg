package signature

import (
	"net/url"
	"testing"
	"time"
)

const (
	key    = "sign"
	secret = "i1ydX9RtHyuJTrw7frcu"
	ttl    = time.Minute * 10
)

func TestSignature_Generate(t *testing.T) {
	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1 c2")

	authorization, date, err := New(key, secret, ttl).Generate(params)
	t.Log("authorization:", authorization)
	t.Log("date:", date)
	t.Log("err:", err)
}

func TestSignature_Verify(t *testing.T) {

	authorization := "sign S4Pwr9Fd08hR4Po+Fy2jdf9aYN1vl2XGwLtPHSNbdx4="
	date := "2023-05-09 15:40:23"

	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1 c2")

	ok, err := New(key, secret, ttl).Verify(authorization, date, params)
	t.Log(ok)
	t.Log(err)
}
