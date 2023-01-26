package stack

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/errs"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/item"
)

type Stack[T any] struct {
	top    *item.Item[T]
	height int
}

func New[T any](value T) *Stack[T] {
	t := item.New(value)
	s := &Stack[T]{
		top:    t,
		height: 1,
	}
	return s
}

func (s Stack[T]) Top() *item.Item[T] {
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

func (s *Stack[T]) Push(value T) {
	i := item.New(value)
	i.SetNext(s.top)
	s.top = i
	s.height++
}

func (s *Stack[T]) Pop() (item *item.Item[T], err error) {
	if s.height == 0 {
		return nil, errs.HeightZero
	}

	poped := s.top
	s.top = s.top.Next()
	s.height--

	return poped, nil
}
