package etl

import "strings"

func Transform(input map[int][]string) map[string]int  {
	newMap :=  map[string]int{}
	for k, v := range input {
		for _, v := range v {
			newMap[strings.ToLower(v)] = k
		}
	}
	return newMap
}