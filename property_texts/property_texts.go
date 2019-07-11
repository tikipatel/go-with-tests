package propertytexts

import "strings"

// RomanNumeral is a struct
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals is an array
var RomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman is a function
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range RomanNumerals {
		for arabic > numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
