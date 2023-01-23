package linkedlists_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/linkedlists"
)

func TestGettersLen1(t *testing.T) {
	l := linkedlists.New(2)

	head := l.Head()
	if head.Value() != 2 {
		t.Errorf("Expected %d but got %v", 2, head.Value())
	}

	tail := l.Tail()
	if tail.Value() != 2 {
		t.Errorf("Expected %d but got %v", 2, tail.Value())
	}

	length := l.Length()
	if length != 1 {
		t.Errorf("Expected %d but got %d", 1, length)
	}
}

func TestString(t *testing.T) {
	v := []int{1, 2}
	l := linkedlists.New(v)
	str := l.String()
	if !reflect.DeepEqual(reflect.TypeOf(str), reflect.TypeOf(*new(string))) {
		t.Errorf("Expected type string, but got %T", str)
	}
}
