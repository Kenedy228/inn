package inn

import (
	"fmt"
	"unicode/utf8"
)

// validateLength проверяет длину ИНН.
//
// Функция сравнивает фактическую длину строки с ожидаемой.
// В случае несовпадения возвращается ошибка ErrInvalidLength
// с дополнительной информацией о длинах.
//
// Используется во внутренних этапах валидации ИНН.
func validateLength(inn string, expectedLength int) error {
	rc := utf8.RuneCountInString(inn)

	if rc != expectedLength {
		return fmt.Errorf("%w: ожидаемая длина %d, действительная длина %d", ErrInvalidLength, expectedLength, rc)
	}

	return nil
}

// validateDigitsOnly проверяет, что строка ИНН состоит только из цифр.
//
// Если строка пуста, возвращается ошибка ErrEmpty.
// Если обнаружен символ, отличный от цифры, возвращается ошибка
// ErrInvalidCharacter с указанием проблемного символа.
//
// Используется перед разбором строки в числовое представление.
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

// validateNotForbidden проверяет, что ИНН не равен запрещённому значению.
//
// Обычно запрещённые значения представляют собой строки,
// состоящие только из нулей. В случае совпадения возвращается
// ошибка ErrForbiddenValue.
func validateNotForbidden(inn, forbidden string) error {
	if inn == forbidden {
		return fmt.Errorf("%w", ErrForbiddenValue)
	}

	return nil
}

// validateChecksumDigit проверяет корректность контрольной цифры ИНН.
//
// Функция вычисляет контрольную цифру на основе базовых цифр и
// весовых коэффициентов, затем сравнивает её с ожидаемым значением.
//
// В случае несовпадения возвращается ошибка ErrInvalidChecksum
// с указанием ожидаемой и фактической контрольной цифры.
func validateChecksumDigit(baseDigits, weights []int, expectedDigit int) error {
	weightedSum := calculateWeightedSum(baseDigits, weights)
	actualDigit := calculateChecksumDigit(weightedSum)

	if actualDigit != expectedDigit {
		return fmt.Errorf("%w: ожидаемая контрольная цифра %d, действительная контрольная цифра %d", ErrInvalidChecksum, expectedDigit, actualDigit)
	}

	return nil
}
