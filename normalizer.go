package inn

import "strings"

func Normalize(code string) string {
	return strings.ReplaceAll(code, " ", "")
}
