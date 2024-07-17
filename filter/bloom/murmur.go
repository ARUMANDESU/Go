package bloom

import murmur3 "github.com/yihleego/murmurhash3"

type MurmurHasher struct {
}

func (m *MurmurHasher) Hash(data []byte) uint {
	murmur := murmur3.New32()
	h := murmur.HashBytes(data)

	return uint(h.AsInt32())
}
