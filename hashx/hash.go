package hashx

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// Hash Calculate the hash of a string
func Hash(str string, h hash.Hash) string {
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5 Calculate the md5 hash of a string
func Md5(str string) string {
	return Hash(str, md5.New())
}

// Sha1 Calculate the sha1 hash of a string
func Sha1(str string) string {
	return Hash(str, sha1.New())
}

// Sha256 Calculate the sha256 hash of a string
func Sha256(str string) string {
	return Hash(str, sha256.New())
}

// Sha512 Calculate the sha512 hash of a string
func Sha512(str string) string {
	return Hash(str, sha512.New())
}
