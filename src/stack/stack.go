package stack

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/node"
)

type Stack[T any] struct {
	top    *node.Node[T]
	height int
}

func New[T any](value T) *Stack[T] {
	t := node.New(value)
	s := &Stack[T]{
		top:    t,
		height: 1,
	}
	return s
}

func (s Stack[T]) Top() *node.Node[T] {
	return s.top
}

func (s Stack[T]) Height() int {
	return s.height
}

func (s Stack[T]) String() string {
	str := ""
	tmp := s.top
	for tmp != nil {
		str += fmt.Sprintf("%v", tmp.Value())
		str += " "
		tmp = tmp.Next()
	}
	return str
}

func (s Stack[T]) Push(value T) {
}
