package sort

import (
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
)

func merge[T constraints.Ordered](s0, s1, s2 *[]T) {
	l1, l2 := len(*s1), len(*s2)
	*s0 = make([]T, l1+l2)

	i, j := 0, 0
	for i < l1 && j < l2 {
		if (*s1)[i] < (*s2)[j] {
			(*s0)[i+j] = (*s1)[i]
			i++
		} else {
			(*s0)[i+j] = (*s2)[j]
			j++
		}
	}

	for i < l1 {
		(*s0)[i+j] = (*s1)[i]
		i++
	}

	for j < l2 {
		(*s0)[i+j] = (*s2)[j]
		j++
	}
}

func MergeSort[T constraints.Ordered](s *[]T) {
	ls := len(*s)
	mid := ls / 2

	if ls <= 1 {
		return
	}

	s1 := (*s)[mid:]
	s2 := (*s)[:mid]

	MergeSort(&s1)
	MergeSort(&s2)
	merge(s, &s1, &s2)
}
