package sum

import "testing"

func TestSum(t *testing.T) {

	assertCorrectSum := func(t *testing.T, got, expected int, given []int) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%d' want '%d given $v'", got, expected, numbers)
		}
	}

	t.Run("sum of 1...5 = 15", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		cd
		assertCorrectSum(t, got, want, numbers)
	})
}
