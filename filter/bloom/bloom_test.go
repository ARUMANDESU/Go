package bloom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	bloomFilter, err := New(FilterConfig{size: 15, hashes: 2})

	assert.NoError(t, err)
	assert.NotNil(t, bloomFilter)
}

func TestNew_Fail(t *testing.T) {
	t.Run("size is less than 1", func(t *testing.T) {
		cases := []int{0, -1, -100}

		for _, size := range cases {
			t.Run(fmt.Sprintf("size: %d", size), func(t *testing.T) {
				bloomFilter, err := New(FilterConfig{size: size, hashes: 2})

				assert.Nil(t, bloomFilter)
				assert.ErrorIs(t, err, ErrInvalidSize)
			})
		}

	})

	t.Run("number of hashes is negative integer", func(t *testing.T) {
		cases := []int{-1, -5, -100}

		for _, hashes := range cases {
			t.Run(fmt.Sprintf("hashes: %d", hashes), func(t *testing.T) {
				bloomFilter, err := New(FilterConfig{size: 2, hashes: hashes})

				assert.Nil(t, bloomFilter)
				assert.ErrorIs(t, err, ErrInvalidHashes)
			})
		}
	})
}
