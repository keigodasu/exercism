package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(p string) Frequency  {
	freq := make(Frequency)

	words := strings.FieldsFunc(p, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	for _, word := range words {
		word = strings.Trim(strings.ToLower(word),"'" )
		freq[word]++
	}
	return freq
}

