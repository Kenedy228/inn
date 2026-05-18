package inn

import "fmt"

func validateDigitsOnly(code string) error {
	runes := []rune(code)

	if len(runes) == 0 {
		return fmt.Errorf("")
	}

	for _, r := range runes {
		if isDigit(r) {
			continue
		}

		return fmt.Errorf("")
	}

	return nil
}
