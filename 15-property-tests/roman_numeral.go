package roman

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var romanCheckPoints = []RomanNumeral{
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

func ConvertToRoman(n uint16) string {

	var r strings.Builder
	for _, numeral := range romanCheckPoints {
		for n >= numeral.Value {
			r.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return r.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range romanCheckPoints {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
