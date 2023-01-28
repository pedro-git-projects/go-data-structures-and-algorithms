package sort

import "github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"

func swap[T constraints.Ordered](s []T, i0, i1 int) {
	tmp := s[i0]
	s[i0] = s[i1]
	s[i1] = tmp
}

func pivot[T constraints.Ordered](s []T, pivotIdx, endIdx int) int {
	swapIdx := pivotIdx

	for i := pivotIdx + 1; i <= endIdx; i++ {
		if s[i] < s[pivotIdx] {
			swapIdx++
			swap(s, swapIdx, i)
		}
	}
	swap(s, pivotIdx, swapIdx)
	return swapIdx
}
