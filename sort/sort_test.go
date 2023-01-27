package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	got := []int{
		4, 3, 2, 1, 5,
	}
	Bubble(got)

	expect := []int{
		1, 2, 3, 4, 5,
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v, but got %v", expect, got)
	}

}

func TestSelection(t *testing.T) {
	got := []int{
		4, 3, 2, 1, 5,
	}
	Selection(got)

	expect := []int{
		1, 2, 3, 4, 5,
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v, but got %v", expect, got)
	}
}

func TestInsertion(t *testing.T) {
	got := []int{
		4, 3, 2, 1, 5,
	}
	Insertion(got)

	expect := []int{
		1, 2, 3, 4, 5,
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v, but got %v", expect, got)
	}
}

func TestVMerge(t *testing.T) {
	got := []int{
		1, 3, 7, 8, 2, 4, 5, 6,
	}

	MergeSort(&got)

	want := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
