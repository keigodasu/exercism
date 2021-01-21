package dna

import (
	"fmt"
)

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, s := range d {
		if _, ok := h[s]; ok {
			h[s] += 1
		} else {
			return nil, fmt.Errorf("couldn't parse DNA due to existing unsupported character: %s", string(s))
		}
	}
	return h, nil
}
