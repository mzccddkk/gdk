# murmur3

Go Implementation of MurmurHash3

More information about these algorithms can be found at:

- [MurmurHash Homepage](https://sites.google.com/site/murmurhash/)
- [Wikipedia Entry on MurmurHash](https://en.wikipedia.org/wiki/MurmurHash)

## Installation

```go
go get github.com/mzccddkk/gdk/murmur3
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/mzccddkk/gdk/murmur3"
)

func main() {
	str := "Hello World"
	var seed uint32 = 123456

	fmt.Println(murmur3.SumA(str))
	// output: 427197390

	fmt.Println(murmur3.SumAWithSeed(str, seed))
	// output: 4125902693

	h1c, h2c, h3c, h4c := murmur3.SumC(str)
	fmt.Println(h1c, h2c, h3c, h4c)
	// output: 2462775914 1499806821 1090519024 188476316

	h1cs, h2cs, h3cs, h4cs := murmur3.SumCWithSeed(str, seed)
	fmt.Println(h1cs, h2cs, h3cs, h4cs)
	// output: 3637336348 2044961496 2622412725 874480687

	h1f, h2f := murmur3.SumF(str)
	fmt.Println(h1f, h2f)
	// output: 1901405986810282715 9504319040210908199

	h1fs, h2fs := murmur3.SumFWithSeed(str, seed)
	fmt.Println(h1fs, h2fs)
	// output: 1440668659686177044 7465135666505861501
}

```