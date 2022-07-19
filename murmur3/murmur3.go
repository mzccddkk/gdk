package murmur3

const (
	c1a uint32 = 0xcc9e2d51
	c2a uint32 = 0x1b873593

	c1c uint32 = 0x239b961b
	c2c uint32 = 0xab0e9789
	c3c uint32 = 0x38b34ae5
	c4c uint32 = 0xa1e38b93

	c1f uint64 = 0x87c37b91114253d5
	c2f uint64 = 0x4cf5ad432745937f
)

func rotl32(x, y uint32) uint32 {
	return (x << y) | (x >> (32 - y))
}

func rotl64(x, y uint64) uint64 {
	return (x << y) | (x >> (64 - y))
}

func fmix32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func fmix64(h uint64) uint64 {
	h ^= h >> 33
	h *= 0xff51afd7ed558ccd
	h ^= h >> 33
	h *= 0xc4ceb9fe1a85ec53
	h ^= h >> 33
	return h
}
