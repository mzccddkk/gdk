package murmur3

import "unsafe"

// SumC Calculate the MurmurHash3 hash of a string with 0 seed
func SumC(str string) (uint32, uint32, uint32, uint32) {
	return SumCWithSeed(str, 0)
}

// SumCWithSeed Calculate the MurmurHash3 hash of a string
func SumCWithSeed(str string, seed uint32) (uint32, uint32, uint32, uint32) {
	h1 := seed
	h2 := seed
	h3 := seed
	h4 := seed

	// body
	b := []byte(str)
	nblocks := len(b) / 16
	for i := 0; i < nblocks; i++ {
		f := (*[4]uint64)(unsafe.Pointer(&b[i*16]))
		k1, k2, k3, k4 := uint32(f[0]), uint32(f[1]), uint32(f[2]), uint32(f[3])

		k1 *= c1c
		k1 = rotl32(k1, 15)
		k1 *= c2c
		h1 ^= k1
		h1 = rotl32(h1, 19)
		h1 += h2
		h1 = h1*5 + 0x561ccd1b

		k2 *= c2c
		k2 = rotl32(k2, 16)
		k2 *= c3c
		h2 ^= k2
		h2 = rotl32(h2, 17)
		h2 += h3
		h2 = h2*5 + 0x0bcaa747

		k3 *= c3c
		k3 = rotl32(k3, 17)
		k3 *= c4c
		h3 ^= k3
		h3 = rotl32(h3, 15)
		h3 += h4
		h3 = h3*5 + 0x96cd1c35

		k4 *= c4c
		k4 = rotl32(k4, 18)
		k4 *= c1c
		h4 ^= k4
		h4 = rotl32(h4, 13)
		h4 += h1
		h4 = h4*5 + 0x32ac3b17
	}

	// tail
	tail := b[nblocks*16:]
	var k1, k2, k3, k4 uint32
	switch len(tail) & 15 {
	case 15:
		k4 ^= uint32(tail[14]) << 16
		fallthrough
	case 14:
		k4 ^= uint32(tail[13]) << 8
		fallthrough
	case 13:
		k4 ^= uint32(tail[12]) << 0
		k4 *= c4c
		k4 = rotl32(k4, 18)
		k4 *= c1c
		h4 ^= k4
		fallthrough

	case 12:
		k3 ^= uint32(tail[11]) << 24
		fallthrough
	case 11:
		k3 ^= uint32(tail[10]) << 16
		fallthrough
	case 10:
		k3 ^= uint32(tail[9]) << 8
		fallthrough
	case 9:
		k3 ^= uint32(tail[8]) << 0
		k3 *= c3c
		k3 = rotl32(k3, 17)
		k3 *= c4c
		h3 ^= k3
		fallthrough

	case 8:
		k2 ^= uint32(tail[7]) << 24
		fallthrough
	case 7:
		k2 ^= uint32(tail[6]) << 16
		fallthrough
	case 6:
		k2 ^= uint32(tail[5]) << 8
		fallthrough
	case 5:
		k2 ^= uint32(tail[4]) << 0
		k2 *= c2c
		k2 = rotl32(k2, 16)
		k2 *= c3c
		h2 ^= k2
		fallthrough

	case 4:
		k1 ^= uint32(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0]) << 0
		k1 *= c1c
		k1 = rotl32(k1, 15)
		k1 *= c2c
		h1 ^= k1
	}

	// finalization
	h1 ^= uint32(len(b))
	h2 ^= uint32(len(b))
	h3 ^= uint32(len(b))
	h4 ^= uint32(len(b))

	h1 += h2
	h1 += h3
	h1 += h4
	h2 += h1
	h3 += h1
	h4 += h1

	h1 = fmix32(h1)
	h2 = fmix32(h2)
	h3 = fmix32(h3)
	h4 = fmix32(h4)

	h1 += h2
	h1 += h3
	h1 += h4
	h2 += h1
	h3 += h1
	h4 += h1

	return h1, h2, h3, h4
}

// SumF Calculate the MurmurHash3 hash of a string with 0 seed
func SumF(str string) (uint64, uint64) {
	return SumFWithSeed(str, 0)
}

// SumFWithSeed Calculate the MurmurHash3 hash of a string
func SumFWithSeed(str string, seed uint32) (uint64, uint64) {
	h1 := uint64(seed)
	h2 := uint64(seed)

	// body
	b := []byte(str)
	nblocks := len(b) / 16
	for i := 0; i < nblocks; i++ {
		f := (*[2]uint64)(unsafe.Pointer(&b[i*16]))
		k1, k2 := f[0], f[1]

		k1 *= c1f
		k1 = rotl64(k1, 31)
		k1 *= c2f
		h1 ^= k1

		h1 = rotl64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		k2 *= c2f
		k2 = rotl64(k2, 33)
		k2 *= c1f
		h2 ^= k2

		h2 = rotl64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	// tail
	tail := b[nblocks*16:]
	var k1, k2 uint64
	switch len(tail) & 15 {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8]) << 0
		k2 *= c2f
		k2 = rotl64(k2, 33)
		k2 *= c1f
		h2 ^= k2
		fallthrough

	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0]) << 0
		k1 *= c1f
		k1 = rotl64(k1, 31)
		k1 *= c2f
		h1 ^= k1
	}

	// finalization
	h1 ^= uint64(len(b))
	h2 ^= uint64(len(b))

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	h2 += h1

	return h1, h2
}
