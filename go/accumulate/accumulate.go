package accumulate

func Accumulate(input []string, fun func(string) string) []string  {
	var list []string

	for _, s := range input {
		list = append(list, fun(s))
	}

	return list
}
