package sort

import "github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"

func Insertion[T constraints.Ordered](s []T) {
	for i := range s {
		tmp := s[i]
		j := i - 1

		for j >= 0 && tmp < s[j] {
			s[j+1] = s[j]
			s[j] = tmp
			j--
		}
	}
}
