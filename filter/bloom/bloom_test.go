package bloom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFilter_LookUp(t *testing.T) {
	tests := []struct {
		name     string
		insert   []string
		find     []string
		expected []bool
	}{
		{
			name:     "simple",
			insert:   []string{"hello", "world"},
			find:     []string{"hello", "world"},
			expected: []bool{true, true},
		},
		{
			name:     "not found",
			insert:   []string{"hello", "world"},
			find:     []string{"hello", "world", "stio'desroitedori'stedeios'"},
			expected: []bool{true, true, false},
		},
		{
			name:     "empty",
			insert:   []string{},
			find:     []string{"hello", "world", "golang"},
			expected: []bool{false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomFilter, err := New(FilterConfig{size: 50, hashes: 2})
			require.NoError(t, err)

			for _, data := range tt.insert {
				bloomFilter.Insert([]byte(data))
			}

			for i, data := range tt.find {
				assert.Equal(t, tt.expected[i], bloomFilter.LookUp([]byte(data)), "data: %s", data)
			}
		})
	}

}
