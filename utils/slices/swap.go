package slices

import "github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"

func Swap[T constraints.Ordered](s []T, i0, i1 int) {
	tmp := s[i0]
	s[i0] = s[i1]
	s[i1] = tmp
}
