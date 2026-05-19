package inn

// calculateWeightedSum вычисляет взвешенную сумму цифр ИНН.
//
// Каждая цифра умножается на соответствующий весовой коэффициент,
// после чего результаты суммируются.
//
// Например:
//
//	digits  = [1, 2, 3]
//	weights = [3, 2, 1]
//	результат = 1*3 + 2*2 + 3*1 = 10
//
// Функция паникует, если длины срезов digits и weights не совпадают,
// так как это считается ошибкой программиста.
func calculateWeightedSum(digits, weights []int) int {
	if len(digits) != len(weights) {
		panic("количество цифр и весовых коэффициентов не совпадает")
	}

	sum := 0

	for i := range digits {
		sum += digits[i] * weights[i]
	}

	return sum
}

// calculateChecksumDigit вычисляет контрольную цифру ИНН
// на основе взвешенной суммы.
//
// Алгоритм соответствует правилам ФНС:
//
//  1. Вычисляется остаток от деления weightedSum на 11.
//  2. Если остаток больше 9, берётся остаток от деления на 10.
//
// Таким образом:
//
//	remainder = weightedSum % 11
//	if remainder > 9 {
//	    remainder = remainder % 10
//	}
//
// Полученное значение является контрольной цифрой.
func calculateChecksumDigit(weightedSum int) int {
	remainder := weightedSum % 11

	if remainder > 9 {
		remainder %= 10
	}

	return remainder
}
