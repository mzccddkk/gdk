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
	fmt.Println(murmur3.SumA("Hello World"))
	// output: 427197390

	fmt.Println(murmur3.SumAWithSeed("Hello World", 123456))
	// output: 4125902693
}

```