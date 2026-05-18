package inn

import (
	"fmt"
	"unicode/utf8"
)

func validateCodeLength(code string, expectedLength int) error {
	rc := utf8.RuneCountInString(code)

	if rc != expectedLength {
		return fmt.Errorf("")
	}

	return nil
}

func validateDigitsOnly(code string) error {
	runes := []rune(code)

	if len(runes) == 0 {
		return fmt.Errorf("")
	}

	for _, r := range runes {
		if isDigit(r) {
			continue
		}

		return fmt.Errorf("")
	}

	return nil
}

func validateChecksumDigit(baseDigits, weights []int, expectedDigit int) error {
	weightedSum := calculateWeightedSum(baseDigits, weights)
	actualDigit := calculateChecksumDigit(weightedSum)

	if actualDigit != expectedDigit {
		return fmt.Errorf("")
	}

	return nil
}
