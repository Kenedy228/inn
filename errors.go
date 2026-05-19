package inn

import "errors"

var (
	ErrEmpty            = errors.New("ИНН не может быть пустым")
	ErrInvalidLength    = errors.New("некорректная длина ИНН")
	ErrInvalidCharacter = errors.New("ИНН должен содержать только цифры")
	ErrForbiddenValue   = errors.New("недопустимое значение ИНН")
	ErrInvalidChecksum  = errors.New("некорректная контрольная сумма ИНН")
)
