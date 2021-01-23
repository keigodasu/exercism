package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	lowerText := strings.ToLower(input)
	for i, c := range lowerText {
		if unicode.IsLetter(c) && strings.ContainsRune(lowerText[i+1:], c) {
			return false
		}
	}

	return true
}
