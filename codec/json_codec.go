package codec

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
)

// JSONEncoding json格式
type JSONEncoding struct{}

// Marshal json encode
func (j JSONEncoding) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal json decode
func (j JSONEncoding) Unmarshal(data []byte, value any) error {
	return json.Unmarshal(data, value)
}

// JSONGzipEncoding json and gzip
type JSONGzipEncoding struct{}

// Marshal json encode and gzip
func (jz JSONGzipEncoding) Marshal(v any) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	// var bufSizeBefore = len(buf)

	buf, err = GzipEncode(buf)
	// log.Infof("gzip_json_compress_ratio=%d/%d=%.2f", bufSizeBefore, len(buf), float64(bufSizeBefore)/float64(len(buf)))
	return buf, err
}

// Unmarshal json encode and gzip
func (jz JSONGzipEncoding) Unmarshal(data []byte, value any) error {
	jsonData, err := GzipDecode(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, value)
	if err != nil {
		return err
	}
	return nil
}

// GzipEncode 编码
func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(in)
	if err != nil {
		err = writer.Close()
		if err != nil {
			return out, err
		}
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

// GzipDecode 解码
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer func() {
		err = reader.Close()
		if err != nil {
			fmt.Printf("reader close err: %+v", err)
		}
	}()

	return io.ReadAll(reader)
}
