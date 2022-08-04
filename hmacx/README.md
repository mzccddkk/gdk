# hmacx

Go hmac functions

## Installation

```go
go get github.com/mzccddkk/gdk/hmacx
```

## Functions

hmac:

```go
func Hmac(str string, key string, h func() hash.Hash) string

func Md5(str string, key string) string

func Sha1(str string, key string) string

func Sha256(str string, key string) string

func Sha512(str string, key string) string
```

hmacfile:

```go
func HmacFile(filename string, key string, h func() hash.Hash) (string, error)

func Md5File(filename string, key string) (string, error)

func Sha1File(filename string, key string) (string, error)

func Sha256File(filename string, key string) (string, error)

func Sha512File(filename string, key string) (string, error)
```
