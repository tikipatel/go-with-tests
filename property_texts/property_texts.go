package propertytexts

import "strings"

// RomanNumeral is a struct
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals is an array
var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
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
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// ConvertToArabic is a func
func ConvertToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}
