package stringset

import (
	"fmt"
)

type Set map[string]bool

func New() Set  {
	return Set{}
}

func NewFromSlice(arr []string) Set {
	set := New()
	for _, v := range arr {
		set.Add(v)
	}

	return set
}

func (s Set) String() string {
	retString := "{"
	for k, _ := range s {
		if retString != "{" {
			retString += ", "
		}
		retString += fmt.Sprintf("%q", k)
	}
	retString += "}"

	return retString
}

func (s Set) IsEmpty() bool{
	return len(s) == 0
}

func (s Set) Has(e string) bool{
	_, exist := s[e]
	return exist
}

func Subset(s1, s2 Set) bool  {
	for k, _ := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool  {
	for k, _ := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool  {
	if len(s1) != len(s2) {
		return false
	}

	for k, _ := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func (s Set) Add(e string)  {
	s[e] = true
}

func Intersection(s1, s2 Set) Set  {
	s := New()
	for k, _ := range s1 {
		if s2.Has(k) {
			s.Add(k)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set  {
	s := New()
	for k, _ := range s1{
		if !s2.Has(k) {
			s.Add(k)
		}
	}
	return s
}

func Union(s1, s2 Set) Set  {
	for k, _ := range s1 {
		s2.Add(k)
	}

	return s2
}

