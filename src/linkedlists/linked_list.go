package linkedlists

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/node"
)

type LinkedList[T any] struct {
	head   *node.Node[T]
	tail   *node.Node[T]
	length int
}

func New[T any](value T) *LinkedList[T] {
	n := node.New(value)
	ll := &LinkedList[T]{
		head:   n,
		tail:   n,
		length: 1,
	}
	return ll
}

func (ll LinkedList[T]) String() string {
	temp := ll.head
	acc := ""
	for temp != nil {
		acc += fmt.Sprintf("%v", temp.Value())
		acc += " "
		temp = temp.Next()
	}
	return acc
}

func (ll LinkedList[T]) Head() *node.Node[T] {
	return ll.head
}

func (ll LinkedList[T]) Tail() *node.Node[T] {
	return ll.tail
}

func (ll LinkedList[T]) Length() int {
	return ll.length
}
