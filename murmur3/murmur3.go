package murmur3

const (
	c132 uint32 = 0xcc9e2d51
	c232 uint32 = 0x1b873593

	c1128 uint32 = 0x239b961b
	c2128 uint32 = 0xab0e9789
	c3128 uint32 = 0x38b34ae5
	c4128 uint32 = 0xa1e38b93
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
