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
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		assertCorrectPerimeter(t, got, want)
	})

	t.Run("Perimeter of rect with height 10 and width of 20 should be 40", func(t *testing.T) {
		rectangle := Rectangle{20, 10}
		got := Perimeter(rectangle)
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
		rectangle := Rectangle{11.0, 34.0}
		got := rectangle.Area()
		want := 374.0

		assertCorrectArea(t, got, want)
	})

	t.Run("Area of a circle with radius 10 should be 314.1592653589793", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		assertCorrectArea(t, got, want)
	})
}

func TestShapesArea(t *testing.T) {

	assertCorrectArea := func(t *testing.T, name string, actual float64, expected float64) {
		t.Helper()
		if actual != expected {
			t.Errorf("%#vot %.2f want %.2f", name, actual, expected)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertCorrectArea(t, tt.name, got, tt.want)
	}
}
