package dlinkedlist_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/dlinkedlist"
)

func TestGettersLen1(t *testing.T) {
	l := dlinkedlist.New(2)

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
	l := dlinkedlist.New(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	head := l.Head()
	if head.Value() != 2 {
		t.Errorf("Expected %d but got %v", 2, head.Value())
	}

	next := l.Head().Next()
	if next.Value() != 3 {
		t.Errorf("Expected %d but got %v", 3, next.Value())
	}

	tail := l.Tail()
	if tail.Value() != 5 {
		t.Errorf("Expected %d but got %v", 5, tail.Value())
	}

	prev := l.Tail().Prev()
	if prev.Value() != 4 {
		t.Errorf("Expected %d but got %v", 4, prev.Value())
	}

	length := l.Length()
	if length != 4 {
		t.Errorf("Expected %d but got %d", 4, length)
	}
}

func TestString(t *testing.T) {
	v := []int{1, 2}
	l := dlinkedlist.New(v)
	str := l.String()
	if !reflect.DeepEqual(reflect.TypeOf(str), reflect.TypeOf(*new(string))) {
		t.Errorf("Expected type string, but got %T", str)
	}
}

func TestAppend(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	err := l.Append(3)
	expected := 3

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveLastNonNil(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	l.Append(3)
	err := l.RemoveLast()
	expected := 2

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveLastNil(t *testing.T) {
	l := dlinkedlist.New(1)
	err := l.RemoveLast()
	expected := 0

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if err != nil {
		t.Error(err)
	}

	if l.Head() != nil || l.Tail() != nil {
		t.Errorf("Expected %v but got %v and %v", nil, l.Head(), l.Tail())

	}
}

func TestRemoveOnly(t *testing.T) {
	l := dlinkedlist.New(1)
	err := l.RemoveFirst()
	expected := 0

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	err := l.RemoveFirst()
	expected := 1

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	l.Prepend(0)
	l.Append(3)

	result, _ := l.Get(0)
	expected := 0

	if result.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, result.Value())
	}
}

func TestSet(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	l.Set(1, 0)
	got, _ := l.Get(1)
	expected := 0

	if got.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, got.Value())
	}

	expected = 2
	length := l.Length()

	if length != expected {
		t.Errorf("Expected %d but got %d", expected, length)
	}
}

func TestInsert(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	l.Prepend(0)
	l.Append(3)

	l.Insert(2, 9)
	v, _ := l.Get(2)
	got := v.Value()

	expected := 9

	if got != expected {
		t.Errorf("Expected %d but got %d", expected, got)
	}

	expected = 5
	length := l.Length()

	if length != expected {
		t.Errorf("Expected %d but got %d", expected, length)
	}
}

func TestRemove(t *testing.T) {
	l := dlinkedlist.New(1)
	l.Append(2)
	l.Append(3)
	b4, _ := l.Get(3)
	l.Append(4)
	l.Prepend(0)

	err := l.Remove(3)
	after, _ := l.Get(3)

	if b4 == after {
		t.Errorf("Expected %v but got %v", b4, after)
	}

	if err != nil {
		t.Error(err)
	}
}
