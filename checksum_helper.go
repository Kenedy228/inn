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
