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

func TestMerge(t *testing.T) {
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

func TestSwap(t *testing.T) {
	got := []int{
		3, 4,
	}

	swap(got, 0, 1)

	want := []int{
		4, 3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func TestPivot(t *testing.T) {
	got := []int{
		4, 6, 1, 7, 3, 2, 5,
	}

	gotIdx := pivot(got, 0, len(got)-1)
	wantIdx := 3

	want := []int{
		2, 1, 3, 4, 6, 7, 5,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}

	if !reflect.DeepEqual(gotIdx, wantIdx) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
