package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func HmacSha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func HmacSha1(secretToken, payloadBody string) []byte {
	mac := hmac.New(sha1.New, []byte(secretToken))
	mac.Write([]byte(payloadBody))
	return mac.Sum(nil)
}
