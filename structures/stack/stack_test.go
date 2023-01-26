package stack_test

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/stack"
)

func TestPush(t *testing.T) {
	s := stack.New(1)
	s.Push(2)

	expect := 2
	got := s.Height()

	if expect != got {
		t.Errorf("Expected %d but got %d", expect, got)
	}
}

func TestPop(t *testing.T) {
	s := stack.New(1)
	s.Push(2)
	s.Push(3)

	popped, err := s.Pop()
	expected := 3

	if popped.Value() != expected {
		t.Errorf("Expected %d but got %d", expected, popped.Value())
	}

	if err != nil {
		t.Error(err)
	}

}
