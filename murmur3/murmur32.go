package murmur3

// SumA Calculate the MurmurHash3 hash of a string with 0 seed
func SumA(str string) uint32 {
	return SumAWithSeed(str, 0)
}

// SumAWithSeed Calculate the MurmurHash3 hash of a string
func SumAWithSeed(str string, seed uint32) uint32 {
	h := seed

	// body
	b := []byte(str)
	l := len(b) / 4 * 4
	for i := 0; i < l; i += 4 {
		f := b[i : i+4]
		k := uint32(f[0]) | uint32(f[1])<<8 | uint32(f[2])<<16 | uint32(f[3])<<24

		k *= c1a
		k = rotl32(k, 15)
		k *= c2a

		h ^= k
		h = rotl32(h, 13)
		h = h*5 + 0xe6546b64
	}

	// tail
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
		k1 *= c1a
		k1 = rotl32(k1, 15)
		k1 *= c2a
		h ^= k1
	}

	// finalization
	h ^= uint32(len(b))

	return fmix32(h)
}
