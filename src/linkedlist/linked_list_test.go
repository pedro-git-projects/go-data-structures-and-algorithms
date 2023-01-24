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
	l.Append(2)
	l.Append(3)
	expected := 3

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}
}

func TestRemoveLastNonNil(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Append(3)
	l.RemoveLast()
	expected := 2

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

}

func TestRemoveLastNil(t *testing.T) {
	l := linkedlist.New(1)
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

func TestPrepend(t *testing.T) {
	l := linkedlist.New(1)
	l.Prepend(0)
	expected := 2

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}
}

func TestPrependEmpty(t *testing.T) {
	l := linkedlist.New(1)
	l.RemoveLast()
	l.Prepend(0)
	expected := 1

	if l.Length() != expected {
		t.Errorf("Expected %d but got %d", expected, l.Length())
	}

	_, err := l.GetByIndex(0)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveOnly(t *testing.T) {
	l := linkedlist.New(1)
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
	l := linkedlist.New(1)
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

func TestGetByIndex(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Prepend(0)
	l.Append(3)

	result, _ := l.GetByIndex(0)
	expected := 0

	if result.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, result.Value())
	}

	result, _ = l.GetByIndex(2)
	expected = 2

	if result.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, result.Value())
	}
}

func TestSet(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Set(1, 0)
	got, _ := l.GetByIndex(1)
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

func TestRemove(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Append(3)
	b4, _ := l.GetByIndex(3)
	l.Append(4)
	l.Prepend(0)

	err := l.Remove(3)
	after, _ := l.GetByIndex(3)

	if b4 == after {
		t.Errorf("Expected %v but got %v", b4, after)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Prepend(0)
	l.Append(3)

	l.Insert(2, 9)
	val, _ := l.GetByIndex(2)
	got := val.Value()

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

func TestReverse(t *testing.T) {
	l := linkedlist.New(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.Append(7)

	var original []int
	for i := 0; i < l.Length(); i++ {
		n, _ := l.GetByIndex(i)
		original = append(original, n.Value())
	}

	l.Reverse()
	var reversed []int
	for i := 0; i < l.Length(); i++ {
		n, _ := l.GetByIndex(i)
		reversed = append(reversed, n.Value())
	}

	expected := []int{
		7, 6, 5, 4, 3, 2, 1,
	}

	if !reflect.DeepEqual(expected, reversed) {
		t.Errorf("Expected %v but got %v", expected, reversed)
	}

	if reflect.DeepEqual(original, reversed) {
		t.Errorf("Expected %v but got %v", expected, reversed)
	}
}
