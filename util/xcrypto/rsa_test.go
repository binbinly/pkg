package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	PublicPemFile       = "/test/testdata/rsa_public.pem"
	PublicPKCS1PemFile  = "/test/testdata/rsa_public_pkcs1.pem"
	PrivatePemFile      = "/test/testdata/rsa_private.pem"
	PrivatePKCS1PemFile = "/test/testdata/rsa_private_pkcs1.pem"
)

var (
	rootDir string
	key     []byte
	origin  []byte
	iv      []byte
)

func TestMain(m *testing.M) {
	//当前目录
	currDir, _ := os.Getwd()
	rootDir = filepath.Dir(filepath.Dir(currDir))
	// 声明一个16字节的key
	key = []byte("example key 1234")
	// 声明一个随意长度的 需加密内容
	origin = []byte("need to crypto encode test text")
	// 声明一个16字节的iv
	iv = []byte("example iv tests")
	m.Run()
}

// 生成RSA密钥对
func TestGenerateRSAKey(t *testing.T) {
	// 声明位数
	var bits = 1024

	// 使用随机数据生成器random生成一对具有指定位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	// 使用指定的位数生成一对多质数的RSA密钥。(质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数)
	// 虽然公钥可以和二质数情况下的公钥兼容（事实上，不能区分两种公钥），私钥却不行。
	// 因此有可能无法生成特定格式的多质数的密钥对，或不能将生成的密钥用在其他（语言的）代码里
	//privateKey, err := rsa.GenerateMultiPrimeKey(rand.Reader, 5, bits)

	if err != nil {
		t.Fatal(err)
	}

	// 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derPrivate := x509.MarshalPKCS1PrivateKey(privateKey)
	// 初始化一个PEM编码的结构
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivate,
	}
	// 创建文件，如果文件存在内容重置为空
	file, err := os.Create(rootDir + PrivatePKCS1PemFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// 将Block的pem编码写入文件
	err = pem.Encode(file, priBlock)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("RSA-PKCS1私钥生成成功")

	// 将rsa私钥序列化为PKCS#8 DER编码
	derPrivate8, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// 初始化一个PEM编码的结构
	priBlock8 := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivate8,
	}

	// 创建文件，如果文件存在内容重置为空
	file2, err := os.Create(rootDir + PrivatePemFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	// 将Block的pem编码写入文件
	err = pem.Encode(file2, priBlock8)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("RSA-PKCS8私钥生成成功")

	// 获取公钥
	publicKey := &privateKey.PublicKey

	// 将公钥序列化为PKIX格式DER编码
	derPublic, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	// 初始化一个PEM编码的结构
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic,
	}

	// 创建文件，如果文件存在内容重置为空
	file3, err := os.Create(rootDir + PublicPemFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file3.Close()

	// 将Block的pem编码写入文件
	err = pem.Encode(file3, pubBlock)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("RSA-PKIX公钥生成成功")

	derPublic1 := x509.MarshalPKCS1PublicKey(publicKey)
	// 初始化一个PEM编码的结构
	pubBlock1 := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic1,
	}
	// 创建文件，如果文件存在内容重置为空
	file4, err := os.Create(rootDir + PublicPKCS1PemFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file4.Close()

	// 将Block的pem编码写入文件
	err = pem.Encode(file4, pubBlock1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("RSA-PKCS1公钥生成成功")
}

// 使用PKCS#1v1.5规定的填充方案和RSA公钥加密/私钥解密
func TestRSAPKCS1v15(t *testing.T) {
	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		t.Fatal(err)
	}

	// rsa公钥加密
	cipherText, err := RsaEncrypt(publicKey, origin)
	if err != nil {
		t.Fatal(err)
	}

	// rsa私钥解密
	originText, err := RsaDecrypt(privateKey, cipherText)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, originText, origin)
}

// RSA-OAEP算法公钥加密/私钥解密
func TestRSAOAEP(t *testing.T) {
	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		t.Fatal(err)
	}

	// 声明label
	var label = []byte("test")

	// rsa公钥加密
	cipherText, err := RsaEncryptOAEP(publicKey, origin, label)
	if err != nil {
		t.Fatal(err)
	}

	// rsa私钥解密
	originText, err := RsaDecryptOAEP(privateKey, cipherText, label)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, originText, origin)
}

// RSA私钥签名/公钥验证
func TestRSASignature(t *testing.T) {
	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		t.Fatal(err)
	}

	// rsa私钥签名
	signature, err := RsaSign(privateKey, origin)
	if err != nil {
		t.Fatal(err)
	}

	//rsa公钥验签
	if err = RsaVerify(publicKey, origin, signature); err != nil {
		t.Fatal(err)
	}
}

// RSA-PASS私钥签名/公钥验证
func TestRSASignPass(t *testing.T) {
	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		t.Fatal(err)
	}

	// 初始化一个 PSS签名 的参数
	var SignOpts = rsa.PSSOptions{SaltLength: 8}

	// 初始化一个 PSS认证 的参数
	var VerifyOpts = rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	// 当PSS签名参数SaltLength 为 rsa.PSSSaltLengthAuto, PSS认证参数SaltLength 必须为 rsa.PSSSaltLengthAuto
	// 当PSS签名参数SaltLength 为 rsa.PSSSaltLengthEqualsHash, PSS认证参数SaltLength 为 rsa.PSSSaltLengthAuto或rsa.PSSSaltLengthEqualsHash
	// 当PSS签名参数SaltLength 为 指定值时 如8, PSS认证参数SaltLength 为rsa.PSSSaltLengthAuto或8
	//

	// rsa私钥签名
	signature, err := RsaSignPass(privateKey, origin, &SignOpts)
	if err != nil {
		t.Fatal(err)
	}

	//rsa公钥验签
	if err = RsaVerifyPass(publicKey, origin, signature, &VerifyOpts); err != nil {
		t.Fatal(err)
	}
}

func getRSAKey() (privateKey []byte, publicKey []byte, err error) {

	// 读取publicKey内容
	publicKey, err = os.ReadFile(rootDir + PublicPemFile)
	if err != nil {
		return
	}
	// 读取privateKey内容
	privateKey, err = os.ReadFile(rootDir + PrivatePemFile)
	return
}
