package xcrypto

import (
	"crypto/rand"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCBC(t *testing.T) {
	cipherText, err := AesCBCEncrypt(origin, key, iv)
	if err != nil {
		t.Fatalf("aes cbc encrypt err:%v", err)
	}

	originText, err := AesCBCDecrypt(cipherText, key, iv)
	if err != nil {
		t.Fatalf("aes cbc decrypt err:%v", err)
	}
	assert.Equal(t, originText, origin)
}

func TestAesCFB(t *testing.T) {

	// 加密
	cipherText, err := AesCFBEncrypt(origin, key, iv)
	if err != nil {
		t.Fatal("aes cfb encrypt err: ", err)
	}

	// 解密
	originText, err := AesCFBDecrypt(cipherText, key, iv)
	if err != nil {
		t.Fatal("aes cfb decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestAesCTR(t *testing.T) {

	// 加密
	cipherText, err := AesCTREncrypt(origin, key, iv)
	if err != nil {
		t.Fatal("aes ctr encrypt err: ", err)
	}

	// 解密
	originText, err := AesCTRDecrypt(cipherText, key, iv)
	if err != nil {
		t.Fatal("aes ctr decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestAesOFB(t *testing.T) {
	// 加密
	cipherText, err := AesOFBEncrypt(origin, key, iv)
	if err != nil {
		t.Fatal("aes ofb encrypt err: ", err)
	}

	// 解密
	originText, err := AesOFBDecrypt(cipherText, key, iv)
	if err != nil {
		t.Fatal("aes ofb decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestAesGCM(t *testing.T) {
	// 初始化一个长度为12字节的空的[]byte，不要使用超过2^32个随机非字符，因为存在重复的风险
	nonce := make([]byte, 12)
	// 使用rand随机生成数据
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	// StreamReader方式加密
	cipherText, err := AesGCMEncrypt(origin, key, nonce)
	if err != nil {
		t.Fatal("aes gcm stream encrypt err: ", err)
	}

	// StreamWriter方式解密
	originText, err := AesGCMDecrypt(cipherText, key, nonce)
	if err != nil {
		t.Fatal("aes gcm stream decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}

func TestAesOFBStream(t *testing.T) {
	// StreamReader方式加密
	cipherText, err := AesOFBEncryptStreamReader(origin, key, iv)
	if err != nil {
		t.Fatal("aes ofb stream encrypt err: ", err)
	}

	// StreamWriter方式解密
	originText, err := AesOFBDecryptStreamWriter(cipherText, key, iv)
	if err != nil {
		t.Fatal("aes ofb stream decrypt err: ", err)
	}
	assert.Equal(t, originText, origin)
}
