# inn

[![Go Version](https://img.shields.io/badge/go-1.26.2-blue.svg)](https://go.dev)
[![CI](https://github.com/Kenedy228/inn/actions/workflows/ci.yml/badge.svg)](https://github.com/Kenedy228/inn/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/Kenedy228/inn.svg)](https://pkg.go.dev/github.com/Kenedy228/inn)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**inn** — минималистичная и быстрая Go библиотека для валидации ИНН РФ.

Поддерживает:
- физические лица (ФЛ);
- индивидуальных предпринимателей (ИП);
- юридические лица (ЮЛ).

Полная проверка включает:
- нормализацию входных данных;
- проверку длины;
- проверку допустимых символов;
- проверку запрещённых значений;
- проверку контрольных цифр по алгоритмам ФНС.

---

## 📦 Установка

```bash
go get github.com/Kenedy228/inn
```

---

## 🚀 Быстрый старт

```go
package main

import (
	"log"

	"github.com/Kenedy228/inn"
)

func main() {
	err := inn.ValidateNaturalPerson("500100732259")
	if err != nil {
		log.Fatal(err)
	}
}
```

---

## 🔍 Использование

### Физическое лицо

```go
err := inn.ValidateNaturalPerson("500100732259")
```

### Индивидуальный предприниматель

```go
err := inn.ValidateIndividualEntrepreneur("500100732259")
```

### Юридическое лицо

```go
err := inn.ValidateLegalEntity("7707083893")
```

---

## 🔧 Нормализация

Перед валидацией строка автоматически нормализуется:

```go
normalized := inn.Normalize("500 100 732 259")
// "500100732259"
```

Удаляются пробелы. Другие символы не изменяются.

---

## ❌ Обработка ошибок

Библиотека возвращает типизированные ошибки:

```go
var (
	ErrEmpty
	ErrInvalidLength
	ErrInvalidCharacter
	ErrForbiddenValue
	ErrInvalidChecksum
)
```

Пример:

```go
err := inn.ValidateNaturalPerson("123")

if err != nil {
	switch {
	case errors.Is(err, inn.ErrInvalidLength):
		// неверная длина
	case errors.Is(err, inn.ErrInvalidChecksum):
		// ошибка контрольной суммы
	}
}
```

Ошибки оборачиваются `%w`, поэтому их можно безопасно проверять через `errors.Is`.

---

## 📚 Алгоритм

Библиотека реализует официальные алгоритмы ФНС:

- ИНН ФЛ и ИП (12 цифр):
  - 2 контрольные цифры.

- ИНН юридического лица (10 цифр):
  - 1 контрольная цифра.

Контрольные цифры рассчитываются через взвешенную сумму и модульные операции (% 11, % 10).

---

## 🧪 Тестирование

```bash
go test ./...
```

---

## 📄 Лицензия

MIT