package xcrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RsaEncryptOAEP 加密
func RsaEncryptOAEP(publicKey, originText, label []byte) ([]byte, error) {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	// 采用RSA-OAEP算法加密指定数据。数据不能超过((公共模数的长度)-2*( hash长度)+2)字节
	// label参数可能包含不加密的任意数据，但这给了信息重要的背景。例如，如果给定公钥用于解密两种类型的消息，然后是不同的标签值可用于确保用于一个目的的密文不能 被攻击者用于另一个目的。如果不需要，可以为空
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, originText, label)
}

// RsaDecryptOAEP 解密
func RsaDecryptOAEP(privateKey, cipherText, label []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。如果random不是nil，函数会注意规避时间侧信道攻击
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, pri, cipherText, label)
}

// RsaSignPass 签名
func RsaSignPass(privateKey, originText []byte, opts *rsa.PSSOptions) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 声明MD5 Hash类型
	hash := crypto.MD5
	// 根据Hash类型创建hash
	h := hash.New()
	// 写入内容
	h.Write(originText)
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	hashed := h.Sum(nil)

	// 采用RSASSA-PSS方案计算签名
	// 注意hashed必须是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果
	// opts参数可以为nil，此时会使用默认参数
	return rsa.SignPSS(rand.Reader, pri, hash, hashed, opts)
}

// RsaVerifyPass 验签
func RsaVerifyPass(publicKey, originText, signature []byte, opts *rsa.PSSOptions) error {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return err
	}

	// 声明MD5 Hash类型
	hash := crypto.MD5
	// 根据Hash类型创建hash
	h := hash.New()
	// 写入内容
	h.Write(originText)
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	hashed := h.Sum(nil)

	// 认证一个PSS签名
	// hashed是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果。合法的签名会返回nil，否则表示签名不合法
	// opts参数可以为nil，此时会使用默认参数
	return rsa.VerifyPSS(pub, hash, hashed, signature, opts)
}

// RsaEncrypt 加密
func RsaEncrypt(publicKey, originText []byte) ([]byte, error) {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法加密msg。信息不能超过((公共模数的长度)-11)字节
	// 注意：使用本函数加密明文（而不是会话密钥）是危险的，请尽量在新协议中使用RSA OAEP
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originText)
}

// RsaDecrypt 解密
func RsaDecrypt(privateKey, cipherText []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。如果random不是nil，函数会注意规避时间侧信道攻击
	return rsa.DecryptPKCS1v15(rand.Reader, pri, cipherText)
}

// RsaSign 签名
func RsaSign(privateKey, originText []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 返回数据的 SHA256 校验和
	hashed := sha256.Sum256(originText)

	// 使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	// 注意hashed必须是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed[:])
}

// RsaVerify 验签
func RsaVerify(publicKey, originText, signature []byte) error {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return err
	}

	// 返回数据的 SHA256 校验和
	hashed := sha256.Sum256(originText)

	// 验证 RSA PKCS＃1 v1.5 签名
	// hashed是使用提供的hash参数对（要签名的）原始数据进行hash的结果。合法的签名会返回nil，否则表示签名不合法
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// BuildRSAPublicKey build PublicKey
func BuildRSAPublicKey(publicKey []byte) (*rsa.PublicKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析一个DER编码的公钥。这些公钥一般在以"BEGIN PUBLIC KEY"出现的PEM块中
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 指定为rsa.PublicKey结构
	pub := pubInterface.(*rsa.PublicKey)

	return pub, nil
}

// BuildRSAPrivateKey build PrivateKey
func BuildRSAPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析一个未加密的PKCS#8私钥
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 指定为rsa.PrivateKey结构
	pri := priInterface.(*rsa.PrivateKey)

	return pri, nil
}

// BuildRSAPKCS1PublicKey build PublicKey
func BuildRSAPKCS1PublicKey(publicKey []byte) (*rsa.PublicKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析一个ASN.1 PKCS#1 DER编码的公钥。
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// BuildRSAPKCS1PrivateKey build PrivateKey
func BuildRSAPKCS1PrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析一个ASN.1 PKCS#1 DER编码的私钥。
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
