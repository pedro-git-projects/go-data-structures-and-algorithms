package sort

import (
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
)

func Selection[T constraints.Ordered](s []T) {
	for i := range s {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if i != min {
			tmp := s[i]
			s[i] = s[min]
			s[min] = tmp
		}
	}

}
