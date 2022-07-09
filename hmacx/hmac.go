package hmacx

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// Hmac Calculate the hash value using the HMAC method
func Hmac(str string, key string, h func() hash.Hash) string {
	hm := hmac.New(h, []byte(key))
	hm.Write([]byte(str))
	return hex.EncodeToString(hm.Sum(nil))
}

// Md5 Calculate the md5 hash value using the HMAC method
func Md5(str string, key string) string {
	return Hmac(str, key, md5.New)
}

// Sha1 Calculate the sha1 hash value using the HMAC method
func Sha1(str string, key string) string {
	return Hmac(str, key, sha1.New)
}

// Sha256 Calculate the sha256 hash value using the HMAC method
func Sha256(str string, key string) string {
	return Hmac(str, key, sha256.New)
}

// Sha512 Calculate the sha512 hash value using the HMAC method
func Sha512(str string, key string) string {
	return Hmac(str, key, sha512.New)
}
