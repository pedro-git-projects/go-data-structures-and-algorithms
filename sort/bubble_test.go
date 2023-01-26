package sort_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/sort"
)

func TestBubble(t *testing.T) {
	got := []int{
		4, 3, 2, 1, 5,
	}
	sort.Bubble(got)

	expect := []int{
		1, 2, 3, 4, 5,
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v, but got %v", expect, got)
	}

}
