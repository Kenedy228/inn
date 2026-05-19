package inn

import (
	"fmt"
	"unicode/utf8"
)

func validateLength(inn string, expectedLength int) error {
	rc := utf8.RuneCountInString(inn)

	if rc != expectedLength {
		return fmt.Errorf("%w: ожидаемая длина %d, действительная длина %d", ErrInvalidLength, expectedLength, rc)
	}

	return nil
}

func validateDigitsOnly(inn string) error {
	runes := []rune(inn)

	if len(runes) == 0 {
		return fmt.Errorf("%w", ErrEmpty)
	}

	for _, r := range runes {
		if isDigit(r) {
			continue
		}

		return fmt.Errorf("%w: запрещенный символ %c", ErrInvalidCharacter, r)
	}

	return nil
}

func validateNotForbidden(inn, forbidden string) error {
	if inn == forbidden {
		return fmt.Errorf("%w", ErrForbiddenValue)
	}

	return nil
}

func validateChecksumDigit(baseDigits, weights []int, expectedDigit int) error {
	weightedSum := calculateWeightedSum(baseDigits, weights)
	actualDigit := calculateChecksumDigit(weightedSum)

	if actualDigit != expectedDigit {
		return fmt.Errorf("%w: ожидаемая контрольная цифра %d, действительная контрольная цифра %d", ErrInvalidChecksum, expectedDigit, actualDigit)
	}

	return nil
}
