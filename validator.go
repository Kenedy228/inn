package inn

import (
	"fmt"
)

func validatePersonINN(inn string) error {
	digits := parseDigits(inn)

	firstPartDigits := digits[:10]
	expectedFirstPartDigit := digits[10]

	if err := validateChecksumDigit(
		firstPartDigits,
		personFirstChecksumWeights,
		expectedFirstPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	secondPartDigits := digits[:11]
	expectedSecondPartDigit := digits[11]

	if err := validateChecksumDigit(
		secondPartDigits,
		personSecondChecksumWeights,
		expectedSecondPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	return nil
}
