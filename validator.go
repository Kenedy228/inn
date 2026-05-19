package inn

import (
	"fmt"
)

func ValidateNaturalPerson(inn string) error {
	inn = Normalize(inn)

	if err := validateDigitsOnly(inn); err != nil {
		return err
	}

	if err := validateLength(inn, naturalPersonDigitsCount); err != nil {
		return err
	}

	if err := validateNotForbidden(inn, naturalPersonForbiddenValue); err != nil {
		return err
	}

	digits := parseDigits(inn)

	firstBase := digits[:10]
	expectedFirstDigit := digits[10]

	if err := validateChecksumDigit(
		firstBase,
		naturalPersonFirstWeights,
		expectedFirstDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	secondDigits := digits[:11]
	expectedSecondDigit := digits[11]

	if err := validateChecksumDigit(
		secondDigits,
		naturalPersonSecondWeights,
		expectedSecondDigit,
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
		legalEntityWeights,
		expectedPartDigit,
	); err != nil {
		return fmt.Errorf("")
	}

	return nil
}
