package inn

import "fmt"

// ValidateNaturalPerson проверяет корректность ИНН физического лица.
//
// Функция нормализует входную строку, проверяет, что ИНН:
//   - не пустой;
//   - состоит только из цифр;
//   - содержит 12 цифр;
//   - не равен запрещённому значению;
//   - имеет корректные первую и вторую контрольные цифры.
//
// Возвращает nil, если ИНН корректен.
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
		return fmt.Errorf("некорректная первая контрольная цифра: %w", err)
	}

	secondDigits := digits[:11]
	expectedSecondDigit := digits[11]

	if err := validateChecksumDigit(
		secondDigits,
		naturalPersonSecondWeights,
		expectedSecondDigit,
	); err != nil {
		return fmt.Errorf("некорректная вторая контрольная цифра: %w", err)
	}

	return nil
}

// ValidateIndividualEntrepreneur проверяет корректность ИНН индивидуального предпринимателя.
//
// Функция нормализует входную строку, проверяет, что ИНН:
//   - не пустой;
//   - состоит только из цифр;
//   - содержит 12 цифр;
//   - не равен запрещённому значению;
//   - имеет корректные первую и вторую контрольные цифры.
//
// Алгоритм проверки контрольных цифр совпадает с алгоритмом для ИНН физического лица.
//
// Возвращает nil, если ИНН корректен.
func ValidateIndividualEntrepreneur(inn string) error {
	inn = Normalize(inn)

	if err := validateDigitsOnly(inn); err != nil {
		return err
	}

	if err := validateLength(inn, individualEntrepreneurDigitsCount); err != nil {
		return err
	}

	if err := validateNotForbidden(inn, individualEntrepreneurForbiddenValue); err != nil {
		return err
	}

	digits := parseDigits(inn)

	firstBase := digits[:10]
	expectedFirstDigit := digits[10]

	if err := validateChecksumDigit(
		firstBase,
		individualEntrepreneurFirstWeights,
		expectedFirstDigit,
	); err != nil {
		return fmt.Errorf("некорректная первая контрольная цифра: %w", err)
	}

	secondDigits := digits[:11]
	expectedSecondDigit := digits[11]

	if err := validateChecksumDigit(
		secondDigits,
		individualEntrepreneurSecondWeights,
		expectedSecondDigit,
	); err != nil {
		return fmt.Errorf("некорректная вторая контрольная цифра: %w", err)
	}

	return nil
}

// ValidateLegalEntity проверяет корректность ИНН юридического лица.
//
// Функция нормализует входную строку, проверяет, что ИНН:
//   - не пустой;
//   - состоит только из цифр;
//   - содержит 10 цифр;
//   - не равен запрещённому значению;
//   - имеет корректную контрольную цифру.
//
// Возвращает nil, если ИНН корректен.
func ValidateLegalEntity(inn string) error {
	inn = Normalize(inn)

	if err := validateDigitsOnly(inn); err != nil {
		return err
	}

	if err := validateLength(inn, legalEntityDigitsCount); err != nil {
		return err
	}

	if err := validateNotForbidden(inn, legalEntityForbiddenValue); err != nil {
		return err
	}

	digits := parseDigits(inn)

	base := digits[:9]
	expectedDigit := digits[9]

	if err := validateChecksumDigit(
		base,
		legalEntityWeights,
		expectedDigit,
	); err != nil {
		return fmt.Errorf("некорректная контрольная цифра: %w", err)
	}

	return nil
}
