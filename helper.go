package inn

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func parseDigits(s string) []int {
	digits := make([]int, 0, len(s))

	for _, r := range s {
		if !isDigit(r) {
			panic("")
		}

		digits = append(digits, int(r-'0'))
	}

	return digits
}
