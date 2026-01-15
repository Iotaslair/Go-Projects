package roman

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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

var ErrOutOfRange = errors.New("roman numerals can't go above 3999")

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", ErrOutOfRange
	}

	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) (uint16, error) {
	var arabic uint16 = 0
	if strings.Contains(roman, "MMMM") {
		return 0, ErrOutOfRange
	}

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic, nil
}
