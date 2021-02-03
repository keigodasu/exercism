package linkedlist

import "errors"

// API:
//
// type Element struct {
//  data int
//  next *Element
// }
//
// type List struct {
//  head *Element
//  size int
// }
//
// func New([]int) *List
// func (*List) Size() int
// func (*List) Push(int)
// func (*List) Pop() (int, error)
// func (*List) Array() []int
// func (*List) Reverse() *List

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(array []int) *List  {
	list := &List{}
	for _, v := range array {
		list.Push(v)
	}
	return list
}

 func (l *List) Size() int {
	 return l.size
 }

 func (l *List) Push(i int){
	 l.head = &Element{
		 data: i,
		 next: l.head,
	 }
	 l.size++
 }


 func (l *List) Pop() (int, error){
 	if l.size < 1 {
		return 0, errors.New("empty")
	}

	head, data := l.head, l.head.data
	l.head = head.next
	head.next  = nil
	l.size--

	 return data, nil
 }

 func (l *List) Array() []int{
 	array := make([]int, l.size)
	 for head, i:=l.head, l.size-1 ; head != nil; head, i = head.next, i-1{
	 	array[i] = head.data
	 }
	 return array
 }

 func (l *List) Reverse() *List{
 	reversedList := &List{}
	 for head := l.head; head != nil; head = head.next {
	 	reversedList.Push(head.data)
	}

	 return reversedList
 }
