package bloom

import "errors"

var (
	ErrInvalidHashes = errors.New("the number of hash functions must be a positive integer")
	ErrInvalidSize   = errors.New("the size of the bit array be more than 0")
)

// FilterConfig is configuration for Bloom Filter
type FilterConfig struct {
	// size is size of bit array
	size int
	// hashes is number of times to hash data
	hashes int
}

// Filter is bloom filter
//
// We cannot delete an element in Bloom Filter.
type Filter struct {
	// bits is bits array
	bits []bool
	// inserted is number of elements that have been inserted
	inserted int
	// hashes is number of times to hash data
	hashes int

	hashers []Hasher
}

type Hasher interface {
	Hash([]byte) uint
}

// New return Bloom Filter instance
func New(cfg FilterConfig) (*Filter, error) {
	if cfg.hashes < 0 {
		return nil, ErrInvalidHashes
	}
	if cfg.size <= 0 {
		return nil, ErrInvalidSize
	}
	hashers := []Hasher{&MurmurHasher{}, &FNVHasher{}}

	if cfg.hashes > len(hashers) {
		return nil, errors.New("number of hashes is too larger")
	}

	return &Filter{
		bits:    make([]bool, cfg.size),
		hashes:  cfg.hashes,
		hashers: hashers,
	}, nil
}

func (f *Filter) Insert(data []byte) {
	for i := 0; i < f.hashes; i++ {
		hash := f.hashers[i].Hash(data)
		f.bits[hash%uint(len(f.bits))] = true
	}
	f.inserted++
}

func (f *Filter) LookUp(data []byte) bool {
	for i := 0; i < f.hashes; i++ {
		hash := f.hashers[i].Hash(data)

		if f.bits[hash%uint(len(f.bits))] == false {
			return false
		}
	}

	return true
}
