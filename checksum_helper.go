package inn

func calculateWeightedSum(digits, weights []int) int {
	if len(digits) != len(weights) {
		panic("")
	}

	sum := 0

	for i := range digits {
		sum += digits[i] * weights[i]
	}

	return sum
}

func calculateChecksumDigit(weightedSum int) int {
	remainder := weightedSum % 11

	if remainder > 9 {
		remainder %= 10
	}

	return remainder
}
