package propertytexts

import "strings"

// RomanNumeral is a struct
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals type
type RomanNumerals []RomanNumeral

// RomanNumerals is an array
var romanNumerals = RomanNumerals{
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

// ValueOf is something
func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range romanNumerals {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

// ConvertToRoman is a function
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range romanNumerals {
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

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10
		// (only valid subtractors)
		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			// get the value of the two character string
			value := romanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total
}
