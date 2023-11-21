package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add 2 and 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got'%d'", expected, sum)
		}
	})
	t.Run("add 3 and 2", func(t *testing.T) {
		sum := Add(3, 2)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got'%d'", expected, sum)
		}
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
