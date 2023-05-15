package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesCBC(t *testing.T) {
	cipherText, err := DesCBCEncrypt(origin, key[:8], iv[:8])
	if err != nil {
		t.Fatalf("des cbc encrypt err:%v", err)
	}

	originText, err := DesCBCDecrypt(cipherText, key[:8], iv[:8])
	if err != nil {
		t.Fatalf("des cbc decrypt err:%v", err)
	}
	assert.Equal(t, originText, origin)
}

func TestDesCFB(t *testing.T) {

	// 加密
	cipherText, err := DesCFBEncrypt(origin, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des cfb encrypt err: ", err)
	}

	// 解密
	originText, err := DesCFBDecrypt(cipherText, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des cfb decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestDesCTR(t *testing.T) {

	// 加密
	cipherText, err := DesCTREncrypt(origin, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ctr encrypt err: ", err)
	}

	// 解密
	originText, err := DesCTRDecrypt(cipherText, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ctr decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestDesOFB(t *testing.T) {
	// 加密
	cipherText, err := DesOFBEncrypt(origin, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ofb encrypt err: ", err)
	}

	// 解密
	originText, err := DesOFBDecrypt(cipherText, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ofb decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestDesOFBStream(t *testing.T) {
	// StreamReader方式加密
	cipherText, err := DesOFBEncryptStreamReader(origin, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ofb stream encrypt err: ", err)
	}

	// StreamWriter方式解密
	originText, err := DesOFBDecryptStreamWriter(cipherText, key[:8], iv[:8])
	if err != nil {
		t.Fatal("des ofb stream decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}
