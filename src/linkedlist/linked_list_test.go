package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/linkedlist"
)

func TestGettersLen1(t *testing.T) {
	l := linkedlist.New(2)

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

func TestGettersLenGte1(t *testing.T) {
	l := linkedlist.New(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	head := l.Head()
	if head.Value() != 2 {
		t.Errorf("Expected %d but got %v", 2, head.Value())
	}

	tail := l.Tail()
	if tail.Value() != 5 {
		t.Errorf("Expected %d but got %v", 5, tail.Value())
	}

	length := l.Length()
	if length != 4 {
		t.Errorf("Expected %d but got %d", 4, length)
	}
}

func TestString(t *testing.T) {
	v := []int{1, 2}
	l := linkedlist.New(v)
	str := l.String()
	if !reflect.DeepEqual(reflect.TypeOf(str), reflect.TypeOf(*new(string))) {
		t.Errorf("Expected type string, but got %T", str)
	}
}

func TestAppend(t *testing.T) {
	l := linkedlist.New(1)
	app := l.Append(2)
	app = l.Append(3)
	expected := 3

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !app {
		t.Errorf("Expected %t but got %t", true, false)
	}

}

func TestRemoveLastNonNil(t *testing.T) {
	l := linkedlist.New(1)
	rem := l.Append(2)
	rem = l.Append(3)
	rem = l.RemoveLast()
	expected := 2

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !rem {
		t.Errorf("Expected %t but got %t", true, false)
	}
}

func TestRemoveLastNil(t *testing.T) {
	l := linkedlist.New(1)
	rem := l.RemoveLast()
	expected := 0

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !rem {
		t.Errorf("Expected %t but got %t", true, false)
	}

	if l.Head() != nil || l.Tail() != nil {
		t.Errorf("Expected %v but got %v and %v", nil, l.Head(), l.Tail())

	}
}

func TestPrepend(t *testing.T) {
	l := linkedlist.New(1)
	b := l.Prepend(0)
	expected := 2

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !b {
		t.Errorf("Expected %t but got %t", true, false)
	}

}

func TestPrependEmpty(t *testing.T) {
	l := linkedlist.New(1)
	l.RemoveLast()
	b := l.Prepend(0)
	expected := 1

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !b {
		t.Errorf("Expected %t but got %t", true, false)
	}

	if l.GetByIndex(0).Next() != nil {
		t.Errorf("Expected %v but got %v", nil, l.GetByIndex(0).Next())
	}

	if l.GetByIndex(0).Value() != 0 {
		t.Errorf("Expected %v but got %v", 0, l.GetByIndex(0).Value())
	}

}

func TestRemoveOnly(t *testing.T) {
	l := linkedlist.New(1)
	b := l.RemoveFirst()
	expected := 0

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !b {
		t.Errorf("Expected %t but got %t", true, false)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	b := l.RemoveFirst()
	expected := 1

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if !b {
		t.Errorf("Expected %t but got %t", true, false)
	}
}

func TestGetByIndex(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Prepend(0)
	l.Append(3)

	result := l.GetByIndex(0)
	expected := 0

	if result.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, result.Value())
	}
}
