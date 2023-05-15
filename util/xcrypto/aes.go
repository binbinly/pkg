package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
)

// AesCBCEncrypt 加密
func AesCBCEncrypt(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回加密字节块的大小
	blockSize := block.BlockSize()

	// PKCS5填充需加密内容
	originText = pKCS5Padding(originText, blockSize)

	// 返回一个密码分组链接模式的、底层用Block加密的cipher.BlockMode，初始向量iv的长度必须等于Block的块尺寸(Block块尺寸等于密钥尺寸)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, len(originText))

	// 加密或解密连续的数据块，将加密内容存储到dst中，src需加密内容的长度必须是块大小的整数倍，src和dst可指向同一内存地址
	blockMode.CryptBlocks(cipherText, originText)

	return cipherText, nil
}

// AesCBCDecrypt 解密
func AesCBCDecrypt(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回一个密码分组链接模式的、底层用b解密的cipher.BlockMode，初始向量iv必须和加密时使用的iv相同
	blockMode := cipher.NewCBCDecrypter(block, iv)

	// 根据 密文[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	originText := make([]byte, len(cipherText))

	// 加密或解密连续的数据块，将解密内容存储到dst中，src需加密内容的长度必须是块大小的整数倍，src和dst可指向同一内存地址
	blockMode.CryptBlocks(originText, cipherText)

	// PKCS5反填充解密内容
	originText = pKCS5UnPadding(originText)
	return originText, nil
}

// AesCFBEncrypt 加密
func AesCFBEncrypt(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, aes.BlockSize+len(originText))

	// 返回一个密码反馈模式的、底层用block加密的cipher.Stream，初始向量iv的长度必须等于block的块尺寸
	stream := cipher.NewCFBEncrypter(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	// cipherText[:aes.BlockSize]为iv值，所以只写入cipherText后面部分
	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

// AesCFBDecrypt 解密
func AesCFBDecrypt(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	// 只使用cipherText除去iv部分
	cipherText = cipherText[aes.BlockSize:]

	// 返回一个密码反馈模式的、底层用block解密的cipher.Stream，初始向量iv必须和加密时使用的iv相同
	stream := cipher.NewCFBDecrypter(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

// AesCTREncrypt 加密
func AesCTREncrypt(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, aes.BlockSize+len(originText))

	// 返回一个计数器模式的、底层采用block生成key流的cipher.Stream，初始向量iv的长度必须等于block的块尺寸
	stream := cipher.NewCTR(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	// cipherText[:aes.BlockSize]为iv值，所以只写入cipherText后面部分
	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

// AesCTRDecrypt 解密
func AesCTRDecrypt(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 只使用cipherText除去iv部分
	cipherText = cipherText[aes.BlockSize:]

	// 返回一个计数器模式的、底层采用block生成key流的cipher.Stream，初始向量iv的长度必须等于block的块尺寸
	stream := cipher.NewCTR(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

// AesGCMEncrypt 加密
func AesGCMEncrypt(originText, key, nonce []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 函数用迦洛瓦计数器模式包装提供的128位Block接口，并返回cipher.AEAD
	g, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 返回加密结果。认证附加的additionalData，将加密结果添加到dst生成新的加密结果，nonce的长度必须是NonceSize()字节，且对给定的key和时间都是独一无二的
	cipherText := g.Seal(nil, nonce, originText, nil)

	return cipherText, nil
}

// AesGCMDecrypt 解密
func AesGCMDecrypt(cipherText, key, nonce []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 函数用迦洛瓦计数器模式包装提供的128位Block接口，并返回cipher.AEAD
	astc, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 返回解密结果。认证附加的additionalData，将解密结果添加到dst生成新的加密结果，nonce的长度必须是NonceSize()字节，nonce和data都必须和加密时使用的相同
	return astc.Open(nil, nonce, cipherText, nil)
}

// AesOFBEncrypt 加密
func AesOFBEncrypt(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, aes.BlockSize+len(originText))

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	// cipherText[:aes.BlockSize]为iv值，所以只写入cipherText后面部分
	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

// AesOFBDecrypt 解密
func AesOFBDecrypt(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 只使用cipherText除去iv部分
	cipherText = cipherText[aes.BlockSize:]

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

// AesOFBEncryptStreamReader 加密
func AesOFBEncryptStreamReader(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 初始化cipher.StreamReader。将一个cipher.Stream与一个io.Reader关联起来，Read方法会调用XORKeyStream方法来处理获取的所有切片
	reader := &cipher.StreamReader{
		S: stream,
		R: bytes.NewReader(originText),
	}

	return io.ReadAll(reader)
}

// AesOFBDecryptStreamWriter 解密
func AesOFBDecryptStreamWriter(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 声明buffer
	var originText bytes.Buffer

	// 初始化cipher.StreamWriter。将一个cipher.Stream与一个io.Writer接口关联起来，Write方法会调用XORKeyStream方法来处理提供的所有切片
	// 如果Write方法返回的n小于提供的切片的长度，则表示StreamWriter不同步，必须丢弃。StreamWriter没有内建的缓存，不需要调用Close方法去清空缓存
	writer := &cipher.StreamWriter{
		S: stream,
		W: &originText,
	}

	// 把reader内容拷贝到writer, writer会调用write方法写入内容
	if _, err = io.Copy(writer, bytes.NewReader(cipherText)); err != nil {
		return nil, err
	}

	return originText.Bytes(), nil
}
