package hashx

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// HashFile Calculate the hash of a file
func HashFile(filename string, h hash.Hash) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Md5File Calculate the md5 hash of a file
func Md5File(filename string) (string, error) {
	return HashFile(filename, md5.New())
}

// Sha1File Calculate the sha1 hash of a file
func Sha1File(filename string) (string, error) {
	return HashFile(filename, sha1.New())
}

// Sha256File Calculate the sha256 hash of a file
func Sha256File(filename string) (string, error) {
	return HashFile(filename, sha256.New())
}

// Sha512File Calculate the sha512 hash of a file
func Sha512File(filename string) (string, error) {
	return HashFile(filename, sha512.New())
}
