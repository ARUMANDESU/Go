package bloom

import "hash/fnv"

type FNVHasher struct {
}

func (f *FNVHasher) Hash(data []byte) uint {
	hasher := fnv.New32()

	_, _ = hasher.Write(data)

	return uint(hasher.Sum32())
}
