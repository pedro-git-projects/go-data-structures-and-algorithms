package sort

import "github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"

func Bubble[T constraints.Ordered](s []T) {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				tmp := s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
			}
		}
	}
}
