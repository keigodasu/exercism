package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("")
	ErrInvalidBase = errors.New("")
)

func FromCodon(input string) (string, error) {
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(input string) ([]string, error) {
	list := []string{}
	for i := 0; i < len(input); i += 3 {
		p, err := FromCodon(input[i : i+3])
		if err == ErrStop {
			return list, nil
		}
		if err == ErrInvalidBase {
			return list, ErrInvalidBase
		}
		list = append(list, p)
	}
	return list, nil
}
