// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	list := regexp.MustCompile(`[\s_-]`).Split(s, -1)
	acronym := ""
	for _, v := range list {
		if v == "" {
			continue
		}
		acronym += v[0:1]
	}

	return strings.ToUpper(acronym)
}