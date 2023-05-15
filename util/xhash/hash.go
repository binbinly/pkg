package xhash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// Sha256Hex string sha256
func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256([]byte(s)))
}

// Sha256 sha256
func Sha256(b []byte) []byte {
	return Hash(b, sha256.New())
}

// Sha512Hex string sha512
func Sha512Hex(s string) string {
	return hex.EncodeToString(Sha512([]byte(s)))
}

// Sha512 sha512
func Sha512(b []byte) []byte {
	return Hash(b, sha512.New())
}

// Sha1Hex string sha1
func Sha1Hex(s string) string {
	return hex.EncodeToString(Sha1([]byte(s)))
}

// Sha1 sha1
func Sha1(b []byte) []byte {
	return Hash(b, sha1.New())
}

// HmacSHA256Hex string hmac sha256
func HmacSHA256Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA256([]byte(s), []byte(key)))
}

// HmacSHA256 hmac sha256
func HmacSHA256(b, key []byte) []byte {
	return Hmac(b, key, sha256.New)
}

// HmacSHA512Hex string hmac sha512
func HmacSHA512Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA512([]byte(s), []byte(key)))
}

// HmacSHA512 hmac sha512
func HmacSHA512(b, key []byte) []byte {
	return Hmac(b, key, sha512.New)
}

// HmacSHA1Hex string hmac sha1
func HmacSHA1Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA1([]byte(s), []byte(key)))
}

// HmacSHA1 hmac sha1
func HmacSHA1(b, key []byte) []byte {
	return Hmac(b, key, sha1.New)
}

// MD5Hex 字符串 md5
func MD5Hex(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// MD5 计算 md5
func MD5(b []byte) []byte {
	return Hash(b, nil)
}

// MD5File 文件MD5
func MD5File(filename string) (string, error) {
	// check if file exists and is not a directory
	info, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", nil
	}

	// read file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return MD5Reader(file)
}

// MD5Reader 计算 md5
func MD5Reader(r io.Reader) (string, error) {
	m := md5.New()
	if _, err := io.Copy(m, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

// Hash 计算hash
func Hash(b []byte, h hash.Hash) []byte {
	if h == nil {
		h = md5.New()
	}
	h.Reset()
	h.Write(b)

	return h.Sum(nil)
}

// Hmac hmac
func Hmac(b []byte, key []byte, h func() hash.Hash) []byte {
	if h == nil {
		h = md5.New
	}
	mac := hmac.New(h, key)
	mac.Write(b)

	return mac.Sum(nil)
}
