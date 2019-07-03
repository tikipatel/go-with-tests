package arrays

import "testing"

func TestSum(t *testing.T) {

	assertExpectedSum := func(t *testing.T, sum, expected int) {
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("[1,5] sum should be 15", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertExpectedSum(t, got, want)
	})
}
