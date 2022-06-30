package murmur3

const (
	c132 uint32 = 0xcc9e2d51
	c232 uint32 = 0x1b873593
	r132 uint32 = 15
	r232 uint32 = 13
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
		k = (k << r132) | (k >> (32 - r132))
		k *= c232

		h ^= k
		h = (h << r232) | (h >> (32 - r232))
		h = h*5 + 0xe6546b64
	}

	rs := b[l:]
	var r uint32
	switch len(rs) & 3 {
	case 3:
		r ^= uint32(rs[2]) << 16
		fallthrough
	case 2:
		r ^= uint32(rs[1]) << 8
		fallthrough
	case 1:
		r ^= uint32(rs[0])
		r *= c132
		r = (r << r132) | (r >> (32 - r132))
		r *= c232
		h ^= r
	}

	h ^= uint32(len(b))

	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}
