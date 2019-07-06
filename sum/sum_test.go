package sum

import "testing"

func TestSum(t *testing.T) {

	assertCorrectSum := func(t *testing.T, got, expected int, numbers []int) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%d' want '%d given %v'", got, expected, numbers)
		}
	}

	t.Run("sum of 1...5 = 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)
	})

	t.Run("sum of 1...4 = 10", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assertCorrectSum(t, got, want, numbers)
	})

	t.Run("sum of 1...7 = 28", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		assertCorrectSum(t, got, want, numbers)
	})
}
