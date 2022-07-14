package murmur3

const (
	c132 uint32 = 0xcc9e2d51
	c232 uint32 = 0x1b873593
)

// Sum32 Calculate the MurmurHash3 hash of a string with 0 seed
func Sum32(str string) uint32 {
	return Sum32WithSeed(str, 0)
}

// Sum32WithSeed Calculate the MurmurHash3 hash of a string
func Sum32WithSeed(str string, seed uint32) uint32 {
	h := seed

	b := []byte(str)
	l := len(b) / 4 * 4
	for i := 0; i < l; i += 4 {
		f := b[i : i+4]
		k := uint32(f[0]) | uint32(f[1])<<8 | uint32(f[2])<<16 | uint32(f[3])<<24

		k *= c132
		k = rotl32(k, 15)
		k *= c232

		h ^= k
		h = rotl32(h, 13)
		h = h*5 + 0xe6546b64
	}

	tail := b[l:]
	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c132
		k1 = rotl32(k1, 15)
		k1 *= c232
		h ^= k1
	}

	h ^= uint32(len(b))

	return fmix32(h)
}

func rotl32(x, y uint32) uint32 {
	return (x << y) | (x >> (32 - y))
}

func fmix32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
