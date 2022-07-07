package hmacx

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Md5 Calculate the md5 hash value using the HMAC method
func Md5(str string, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 Calculate the sha1 hash value using the HMAC method
func Sha1(str string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256 Calculate the sha256 hash value using the HMAC method
func Sha256(str string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512 Calculate the sha512 hash value using the HMAC method
func Sha512(str string, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
