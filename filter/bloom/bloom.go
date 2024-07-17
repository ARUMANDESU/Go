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
	// hasher is hash function
	hasher Hasher
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

	hasher Hasher
}

type Hasher interface {
	Hash([]byte) int
}

// New return Bloom Filter instance
func New(cfg FilterConfig) (*Filter, error) {
	if cfg.hashes < 0 {
		return nil, ErrInvalidHashes
	}
	if cfg.size <= 0 {
		return nil, ErrInvalidSize
	}

	// If hasher is nil, use MurmurHasher by default
	if cfg.hasher == nil {
		cfg.hasher = &MurmurHasher{}
	}

	return &Filter{
		bits:   make([]bool, cfg.size),
		hashes: cfg.hashes,
		hasher: cfg.hasher,
	}, nil
}

func (f *Filter) Insert(data []byte) {
	for i := 0; i < f.hashes; i++ {
		hash := f.hasher.Hash(data)
		f.bits[hash%len(f.bits)] = true
	}
	f.inserted++
}

func (f *Filter) LookUp(data []byte) bool {
	for i := 0; i < f.hashes; i++ {
		hash := f.hasher.Hash(data)

		if f.bits[hash%len(f.bits)] == false {
			return false
		}
	}

	return true
}
