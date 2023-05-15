package xcrypto

import "bytes"

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	// 求填充长度
	padding := blockSize - len(cipherText)%blockSize
	// 将填充长度 重复 长度padding次，返回填充内容
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 将填充内容 添加到 需加密内容
	return append(cipherText, padText...)
}

func pKCS5UnPadding(text []byte) []byte {
	// 获取 解密内容长度
	length := len(text)
	// 获取反填充长度(只获取最后个byte当做长度，因为填充的时候是重复按照长度填充的)
	paddingSize := int(text[length-1])
	// 截取解密内容中的原文内容
	return text[:(length - paddingSize)]
}
