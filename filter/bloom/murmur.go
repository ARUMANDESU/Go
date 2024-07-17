package bloom

import murmur3 "github.com/yihleego/murmurhash3"

type MurmurHasher struct {
}

func (m *MurmurHasher) Hash(data []byte) int {
	murmur := murmur3.New32()
	h := murmur.HashBytes(data)

	return h.Bits()
}
