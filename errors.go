package inn

import "errors"

// Ошибки, возвращаемые функциями валидации ИНН.
var (
	// ErrEmpty возвращается, если переданная строка ИНН пуста.
	ErrEmpty = errors.New("ИНН не может быть пустым")

	// ErrInvalidLength возвращается, если длина ИНН
	// не соответствует ожидаемой.
	ErrInvalidLength = errors.New("некорректная длина ИНН")

	// ErrInvalidCharacter возвращается, если ИНН содержит
	// символы, отличные от цифр.
	ErrInvalidCharacter = errors.New("ИНН должен содержать только цифры")

	// ErrForbiddenValue возвращается, если ИНН равен
	// заведомо недопустимому значению.
	ErrForbiddenValue = errors.New("недопустимое значение ИНН")

	// ErrInvalidChecksum возвращается, если контрольная цифра
	// (или цифры) ИНН не совпадает с рассчитанной по алгоритму.
	ErrInvalidChecksum = errors.New("некорректная контрольная сумма ИНН")
)
