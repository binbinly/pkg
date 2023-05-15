package archive

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTar(t *testing.T) {
	var files = []*File{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}

	// 写tar文件数据流
	buf, err := TarWrite(files)
	if err != nil {
		t.Fatal(err)
	}

	var filepath = "../../test/test.tar"
	// 自动生成并写入文件
	if err = os.WriteFile(filepath, buf.Bytes(), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	fs, err := TarRead(filepath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, files, fs)
}

func TestZip(t *testing.T) {
	var files = []*File{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}

	// 写tar文件数据流
	buf, err := ZipWrite(files)
	if err != nil {
		t.Fatal(err)
	}

	var filepath = "../../test/test.zip"
	// 自动生成并写入文件
	if err = os.WriteFile(filepath, buf.Bytes(), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	fs, err := ZipRead(filepath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, files, fs)
}
