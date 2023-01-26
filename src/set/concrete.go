package set

import (
	"fmt"
	"strings"
)

type set[T comparable] map[T]struct{}

// Assert concrete type:set adheres to Set interface.
var _ Set[string] = (*set[string])(nil)

func newSet[T comparable]() set[T] {
	return make(set[T])
}

func (s *set[T]) Insert(v T) bool {
	prevLen := len(*s)
	(*s)[v] = struct{}{}
	return prevLen != len(*s)
}

func (s *set[T]) insert(v T) {
	(*s)[v] = struct{}{}
}

func (s *set[T]) Cardinality() int {
	return len(*s)
}

func (s *set[T]) Clear() {
	*s = newSet[T]()
}

func (s *set[T]) Clone() Set[T] {
	clonedSet := make(set[T], s.Cardinality())
	for elem := range *s {
		clonedSet.insert(elem)
	}
	return &clonedSet
}

func (s *set[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := (*s)[val]; !ok {
			return false
		}
	}
	return true
}

func (s *set[T]) contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *set[T]) Difference(other Set[T]) Set[T] {
	o := other.(*set[T])

	diff := newSet[T]()
	for elem := range *s {
		if !o.contains(elem) {
			diff.insert(elem)
		}
	}
	return &diff
}

func (s *set[T]) Equals(other Set[T]) bool {
	o := other.(*set[T])

	if s.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range *s {
		if !o.contains(elem) {
			return false
		}
	}
	return true
}

func (s *set[T]) Intersection(other Set[T]) Set[T] {
	o := other.(*set[T])

	intersection := newSet[T]()

	if s.Cardinality() < other.Cardinality() {
		for elem := range *s {
			if o.contains(elem) {
				intersection.insert(elem)
			}
		}
	} else {
		for elem := range *o {
			if s.contains(elem) {
				intersection.insert(elem)
			}
		}
	}
	return &intersection
}

func (s *set[T]) IsProperSubset(other Set[T]) bool {
	return s.Cardinality() < other.Cardinality() && s.IsSubset(other)
}

func (s *set[T]) IsProperSuperset(other Set[T]) bool {
	return s.Cardinality() > other.Cardinality() && s.IsSuperset(other)
}

func (s *set[T]) IsSubset(other Set[T]) bool {
	o := other.(*set[T])
	if s.Cardinality() > other.Cardinality() {
		return false
	}
	for elem := range *s {
		if !o.contains(elem) {
			return false
		}
	}
	return true
}

func (s *set[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *set[T]) Erase(v T) {
	delete(*s, v)
}

func (s *set[T]) String() string {
	items := make([]string, 0, len(*s))

	for elem := range *s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}

func (s *set[T]) Union(other Set[T]) Set[T] {
	o := other.(*set[T])

	n := s.Cardinality()
	if o.Cardinality() > n {
		n = o.Cardinality()
	}
	unionedSet := make(set[T], n)

	for elem := range *s {
		unionedSet.insert(elem)
	}
	for elem := range *o {
		unionedSet.insert(elem)
	}
	return &unionedSet
}
