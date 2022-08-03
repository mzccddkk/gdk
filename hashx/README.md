# hashx

Go hash functions

## Installation

```go
go get github.com/mzccddkk/gdk/hashx
```

## Functions

hash:

```go
func Hash(str string, h hash.Hash) string

func Md5(str string) string

func Sha1(str string) string

func Sha256(str string) string

func Sha512(str string) string
```

hashfile:

```go
func HashFile(filename string, h hash.Hash) (string, error)

func Md5File(filename string) (string, error)

func Sha1File(filename string) (string, error)

func Sha256File(filename string) (string, error)

func Sha512File(filename string) (string, error)
```

bcrypt:

```go
func Bcrypt(password string) string

func BcryptWithCost(password string, cost int) string

func BcryptVerify(hashedPassword string, password string) bool
```