package propertytexts

import "testing"

func TestRomanNumerals(t *testing.T) {

	assertCorrectNumeral := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	// t.Run("", func(t *testing.T) {
	// 	got := ConvertToRoman(1)
	// 	want := "I"

	// 	assertCorrectNumeral(t, got, want)
	// })

	// t.Run("", func(t *testing.T) {
	// 	got := ConvertToRoman(2)
	// 	want := "II"
	// 	assertCorrectNumeral(t, got, want)
	// })

	// t.Run("", func(t *testing.T) {
	// 	got := ConvertToRoman(3)
	// 	want := "III"
	// 	assertCorrectNumeral(t, got, want)
	// })
}
