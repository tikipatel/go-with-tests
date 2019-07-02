package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertExpectedSum := func(t *testing.T, sum, expected int) {
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("2+2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertExpectedSum(t, sum, expected)
	})

	t.Run("4+2 = 6", func(t *testing.T) {
		sum := Add(4, 2)
		expected := 6

		assertExpectedSum(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
