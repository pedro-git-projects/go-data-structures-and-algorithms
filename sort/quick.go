package sort

import (
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/slices"
)

func pivot[T constraints.Ordered](s []T, pivotIdx, endIdx int) int {
	swapIdx := pivotIdx

	for i := pivotIdx + 1; i <= endIdx; i++ {
		if s[i] < s[pivotIdx] {
			swapIdx++
			slices.Swap(s, swapIdx, i)
		}
	}
	slices.Swap(s, pivotIdx, swapIdx)
	return swapIdx
}

// TODO: Remove the need to pass begin and end
func Quick[T constraints.Ordered](s []T, begin, end int) {
	if begin >= end {
		return
	}

	pivotIdx := pivot(s, begin, end)
	Quick(s, begin, pivotIdx-1)
	Quick(s, pivotIdx+1, end)
}
