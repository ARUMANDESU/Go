package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum numbers from 1 to 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("sum numbers [1, 6, 7, 14 ,6 ,3]", func(t *testing.T) {
		numbers := []int{1, 6, 7, 14, 6, 3}

		got := Sum(numbers)
		want := 37

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSum(t, got, want)
	})

	t.Run("sum one empty slice", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 6, 7, 14, 6, 3}, []int{1, 2, 3, 4, 5})
	}
}
