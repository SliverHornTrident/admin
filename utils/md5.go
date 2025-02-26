package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string, b ...byte) string {
	return Md5Byte([]byte(str), b...)
}

func Md5Byte(bytes []byte, b ...byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(b))
}
