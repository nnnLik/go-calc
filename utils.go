package main

import (
	"errors"
	"regexp"
	"strconv"
)

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var arabicToRomanNumerals = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func isRoman(s string) bool {
	pattern := "^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"

	re := regexp.MustCompile(pattern)

	if s == "" {
		return false
	}

	return re.MatchString(s)

}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func romanToArabic(romanNumeral string) (int, error) {
	var result int
	prevValue := 0

	for i := len(romanNumeral) - 1; i >= 0; i-- {
		value, exists := romanNumerals[rune(romanNumeral[i])]
		if !exists {
			return 0, errors.New("Invalid Roman numeral")
		}
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic < 1 || arabic > 3999 {
		return "", errors.New("Input value is out of the valid range.")
	}

	var result string

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range keys {
		for arabic >= value {
			result += arabicToRomanNumerals[value]
			arabic -= value
		}
	}

	return result, nil
}
