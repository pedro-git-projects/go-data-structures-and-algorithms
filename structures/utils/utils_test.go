package utils_test

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/utils"
)

func TestCommutativeSliceComparison(t *testing.T) {
	// equal
	s0 := []int{1, 2, 3, 4, 5}
	s1 := []int{1, 2, 3, 4, 5}
	got := utils.CommutativeSliceComparison(s0, s1)
	expect := true

	if !got == expect {
		t.Errorf("Expected %t but got %t", expect, got)
	}

	// equal but in different order
	s0 = []int{1, 2, 3, 4, 5}
	s1 = []int{2, 1, 5, 4, 3}
	got = utils.CommutativeSliceComparison(s0, s1)
	expect = true

	if !got == expect {
		t.Errorf("Expected %t but got %t", expect, got)
	}

	// missing an element
	s0 = []int{1, 2, 3, 4, 5}
	s1 = []int{1, 5, 4, 3}
	got = utils.CommutativeSliceComparison(s0, s1)
	expect = false

	if !got == expect {
		t.Errorf("Expected %t but got %t", expect, got)
	}

	// an element in excess
	s0 = []int{1, 2, 3, 4, 5}
	s1 = []int{1, 5, 4, 3, 5, 7}
	got = utils.CommutativeSliceComparison(s0, s1)
	expect = false

	if !got == expect {
		t.Errorf("Expected %t but got %t", expect, got)
	}

	// excess element is repeated
	s0 = []int{1, 2, 3, 4, 5}
	s1 = []int{1, 5, 4, 3, 5, 1}
	got = utils.CommutativeSliceComparison(s0, s1)
	expect = false

	if !got == expect {
		t.Errorf("Expected %t but got %t", expect, got)
	}

}
