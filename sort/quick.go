package sort

import (
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/slices"
)

func quickStep[T constraints.Ordered](s []T, a, b int) {
	if a >= b {
		return
	}

	pivot := s[b]

	l := a
	r := b - 1

	for l <= r {
		for l <= r && pivot >= s[l] {
			l++
		}
		for r >= l && s[r] >= pivot {
			r--
		}
		if l < r {
			slices.Swap(s, l, r)
		}
	}
	slices.Swap(s, l, b)
	quickStep(s, a, l-1)
	quickStep(s, l+1, b)
}

func Quick[T constraints.Ordered](s []T) {
	if len(s) <= 1 {
		return
	}
	quickStep(s, 0, len(s)-1)
}
