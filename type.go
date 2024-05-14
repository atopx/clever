package clever

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"unsafe"

	"github.com/atopx/clever/general"
)

// String tips: 只有在原有bytes确保不会发生变化时可以使用
func String(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// Bytes 利用反射转移字符串数据到bytes
func Bytes(src string) (data []byte) {
	return unsafe.Slice(unsafe.StringData(src), len(src))
}

func Last[T any](s []T) *T {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}

func First[T any](s []T) *T {
	if len(s) == 0 {
		return nil
	}
	return &s[0]
}

func QuickFileMd5(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return general.Empty, err
	}
	return BytesMd5(data), nil
}

func FileMd5(filepath string, bufSize int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return general.Empty, err
	}
	defer file.Close()
	hash := md5.New()
	if bufSize > 0 {
		buf := make([]byte, bufSize)
		_, _ = io.CopyBuffer(hash, file, buf)
	} else {
		_, _ = io.Copy(hash, file)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func StringMd5(value string) string {
	return BytesMd5(Bytes(value))
}

func BytesMd5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

func BytesSha1(data []byte) string {
	sum := sha1.Sum(data)
	return hex.EncodeToString(sum[:])
}

func StringSha1(value string) string {
	return BytesSha1(Bytes(value))
}
