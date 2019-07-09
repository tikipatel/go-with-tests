package structs

import "testing"

func TestPerimeter(t *testing.T) {

	assertCorrectPerimeter := func(t *testing.T, actual, expected float64) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected %d got %d", expected, actual)
		}
	}

	t.Run("Perimeter of rect with height and width of 10 should be 40", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		assertCorrectPerimeter(t, got, want)
	})

	t.Run("Perimeter of rect with height 10 and width of 20 should be 40", func(t *testing.T) {
		got := Perimeter(20.0, 10.0)
		want := 60.0

		assertCorrectPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T) {
	assertCorrectArea := func(t *testing.T, actual, expected float64) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected %d got %d", expected, actual)
		}
	}

	t.Run("Area of rectangle with width 11 and height 34 should be 374", func(t *testing.T) {
		got := Area(11.0, 34.0)
		want := 374.0

		assertCorrectArea(t, got, want)
	})
}
