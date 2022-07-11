package hmacx

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// HmacFile Calculate file hash value using the HMAC method
func HmacFile(filename string, key string, h func() hash.Hash) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hm := hmac.New(h, []byte(key))
	_, err = io.Copy(hm, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hm.Sum(nil)), nil
}

// Md5File Calculate file md5 hash value using the HMAC method
func Md5File(filename string, key string) (string, error) {
	return HmacFile(filename, key, md5.New)
}

// Sha1File Calculate file sha1 hash value using the HMAC method
func Sha1File(filename string, key string) (string, error) {
	return HmacFile(filename, key, sha1.New)
}

// Sha256File Calculate file sha256 hash value using the HMAC method
func Sha256File(filename string, key string) (string, error) {
	return HmacFile(filename, key, sha256.New)
}

// Sha512File Calculate file sha512 hash value using the HMAC method
func Sha512File(filename string, key string) (string, error) {
	return HmacFile(filename, key, sha512.New)
}
