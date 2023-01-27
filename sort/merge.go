package sort

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
)

// TODO:
// Right now neither merge nor merge sort have side effects
// I want to keep this functions as void, however if i pass a pointer parameter
// for MergeSort I can't slice it with :
func merge[T constraints.Ordered](s0, s1, s2 []T) {
	l1, l2 := len(s1), len(s2)
	i, j := 0, 0

	s0 = make([]T, l1+l2)

	for i < l1 && j < l2 {
		if s1[i] < s2[j] {
			s0[i+j] = s1[i]
			i++
		} else {
			s0[i+j] = s2[j]
			j++
		}
	}

	for i < l1 {
		s0[i+j] = s1[i]
		i++
	}

	for j < l2 {
		s0[i+j] = s2[j]
		j++
	}

	fmt.Printf("final value: %v\n", s0)
}

func MergeSort[T constraints.Ordered](s []T) {
	ls := len(s)
	mid := ls / 2

	if ls <= 1 {
		return
	}

	s1 := s[mid:]
	s2 := s[:mid]

	MergeSort(s1)
	MergeSort(s2)
	merge(s, s1, s2)
}
