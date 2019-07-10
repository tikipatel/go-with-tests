package sum

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {

	assertCorrectArray := func(t *testing.T, got, expected []int, numbers ...[]int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got '%d' want '%d given %v'", got, expected, numbers)
		}
	}

	t.Run("{1,2} and {3,6} should return {3,9}", func(t *testing.T) {
		got := All([]int{1, 2}, []int{3, 6})
		want := []int{3, 9}

		assertCorrectArray(t, got, want)
	})
}
