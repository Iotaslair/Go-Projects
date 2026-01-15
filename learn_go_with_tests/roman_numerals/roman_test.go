package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
	arabicErrorCases = []uint16{4001, 5236, 65535, 25648}
	RomanErrorCases  = []string{"MMMMI", "MMMMD", "MMMML", "MMMMMMMMMMMM"}
)

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToRoman(test.Arabic)
			if test.Roman != got {
				t.Errorf("output %q != expected %q", got, test.Roman)
			}
			assertNoError(t, err)
		})
	}
}

func TestConvertFromRomanErrorOutOfRange(t *testing.T) {
	for _, arabic := range arabicErrorCases {
		_, err := ConvertToRoman(arabic)
		assertGotError(t, err, ErrOutOfRange.Error())
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToArabic(test.Roman)
			if test.Arabic != got {
				t.Errorf("output %d != expected %d", got, test.Arabic)
			}
			assertNoError(t, err)
		})
	}
}

func TestConvertFromArabicErrorOutOfRange(t *testing.T) {
	for _, roman := range RomanErrorCases {
		_, err := ConvertToArabic(roman)
		assertGotError(t, err, ErrOutOfRange.Error())
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman, _ := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got error when not expected")
	}
}
func assertGotError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error but didn't get one")
	}
	if err.Error() != want {
		t.Errorf("got %q want %q", err.Error(), want)
	}
}
