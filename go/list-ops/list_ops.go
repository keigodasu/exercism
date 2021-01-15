package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (l IntList) Length() int {
	length := 0
	for _, _ = range l {
		length++
	}

	return length
}

func (l IntList) Append(list IntList) IntList {
	appendedList := make(IntList, l.Length()+list.Length())
	for i, v := range l {
		appendedList[i] = v
	}
	for i, v := range list {
		appendedList[i+(l.Length())] = v
	}

	return appendedList
}

func (l IntList) Filter(fn predFunc) IntList {
	filteredList := make(IntList, l.Length())
	i := 0
	for _, v := range l {
		if fn(v) {
			filteredList[i] = v
			i++
		}
	}

	return filteredList[:i]
}

func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < j; j-- {
		l[i], l[j] = l[j], l[i]
		i++
	}

	return l
}

func (l IntList) Map(fn unaryFunc) IntList {
	for i, v := range l {
		l[i] = fn(v)
	}
	return l
}

func (l IntList) Foldl(fn binFunc, arg int) int {
	for _, v := range l {
		arg = fn(arg, v)
	}

	return arg
}

func (l IntList) Foldr(fn binFunc, arg int) int {
	for _, v := range l.Reverse() {
		arg = fn(v, arg)
	}

	return arg
}

func (l IntList) Concat(intLists []IntList) IntList {
	for _, v := range intLists {
		l = l.Append(v)
	}

	return l
}
