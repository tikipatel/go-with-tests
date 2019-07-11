package propertytexts

import (
	"testing"
)

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
		{"4 gets converted to III (can't repeat more than 3 times", 4, "IV"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertCorrectNumeral(t, got, test.Want)
		})
	}
}
