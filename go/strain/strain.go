package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(fn func(int) bool) Ints  {
	var list []int
	for _, v := range i {
		if fn(v) {
			list = append(list, v)
		}
	}
	return list
}

func (i Ints) Discard(fn func(int) bool) Ints  {
	return i.Keep(func(i int) bool {
		return !fn(i)
	})
}

func (s Strings) Keep(fn func(string) bool) Strings  {
	var list []string

	for _, v := range s {
		if fn(v) {
			list = append(list, v)
		}
	}

	return list
}

func (l Lists) Keep(fn func([]int) bool) Lists  {
	var list [][]int

	for _, v := range l {
		if fn(v) {
			list = append(list, v)
		}
	}

	return list
}
