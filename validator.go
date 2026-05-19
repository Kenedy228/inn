package inn

import (
	"fmt"
)

func ValidateNaturalPersonINN(inn string) error {
	inn = Normalize(inn)

	if err := validateDigitsOnly(inn); err != nil {
		return err
	}

	if err := validateCodeLength(inn, naturalPersonINNDigitsCount); err != nil {
		return err
	}

	if err := validateForbiddenValue(inn, naturalPersonForbiddenINN); err != nil {
		return err
	}

	digits := parseDigits(inn)

	firstPartDigits := digits[:10]
	expectedFirstPartDigit := digits[10]

	if err := validateChecksumDigit(
		firstPartDigits,
		naturalPersonFirstChecksumWeights,
		expectedFirstPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	secondPartDigits := digits[:11]
	expectedSecondPartDigit := digits[11]

	if err := validateChecksumDigit(
		secondPartDigits,
		naturalPersonSecondChecksumWeights,
		expectedSecondPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

func validateLegalEntityINN(inn string) error {
	digits := parseDigits(inn)

	partDigits := digits[:9]
	expectedPartDigit := digits[9]

	if err := validateChecksumDigit(
		partDigits,
		legalEntityChecksumWeights,
		expectedPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	return nil
}
