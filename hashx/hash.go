package hashx

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Md5 Calculate the md5 hash of a string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 Calculate the sha1 hash of a string
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256 Calculate the sha256 hash of a string
func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512 Calculate the sha512 hash of a string
func Sha512(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
