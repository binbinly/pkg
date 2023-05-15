package xfile

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Exist 文件或目录是否存在
func Exist(s string) bool {
	_, err := os.Stat(s)
	return err == nil || os.IsExist(err)
}

// Is 文件是否存在
func Is(s string) bool {
	info, err := os.Stat(s)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir 目录是否存在
func IsDir(s string) bool {
	info, err := os.Stat(s)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// ImageType 图片类型
func ImageType(read io.Reader) (string, error) {
	src, err := io.ReadAll(read)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(src)
	// 必须引入 image/jpeg,gif,png 包 否则会报 image: unknown format
	_, imgType, err := image.Decode(reader)
	if err != nil {
		return "", err
	}
	return imgType, nil
}

// Ext 获取文件后缀名
func Ext(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 || index == len(filename)-1 {
		return ""
	}
	return filename[index+1:]
}

// Mkdir 创建目录
func Mkdir(path string) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}

// Create 创建文件
func Create(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(content.Bytes()); err != nil {
		return err
	}
	return nil
}
