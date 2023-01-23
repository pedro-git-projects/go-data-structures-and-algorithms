package lnkdlist

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

func (ll LinkedList[T]) Head() *node.Node[T] {
	return ll.head
}

func (ll LinkedList[T]) Tail() *node.Node[T] {
	return ll.tail
}

func (ll LinkedList[T]) Length() int {
	return ll.length
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

func (ll *LinkedList[T]) Append(value T) bool {
	n := node.New(value)
	if ll.length == 0 {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.SetNext(n)
		ll.tail = n
	}
	ll.length++
	return true
}

func (ll *LinkedList[T]) RemoveLast() bool {
	if ll.length == 0 {
		return false
	}

	desired := ll.head
	temp := ll.head
	for temp.Next() != nil {
		desired = temp
		temp = temp.Next()
	}

	ll.tail = desired
	ll.tail.SetNext(nil)
	ll.length--

	if ll.length == 0 {
		ll.head = nil
		ll.tail = nil
	}

	return true
}

func (ll *LinkedList[T]) Prepend(value T) bool {
	n := node.New(value)
	if ll.head != nil {
		n.SetNext(ll.head)
		ll.head = n
	} else {
		ll.head = n
		ll.tail = n
	}
	ll.length++
	return true
}

func (ll *LinkedList[T]) RemoveFirst() bool {
	if ll.length == 0 {
		return false
	}

	desired := ll.head
	ll.head = ll.head.Next()
	desired.SetNext(nil)
	ll.length--

	if ll.length == 0 {
		ll.tail = nil
	}
	return true
}

func (ll *LinkedList[T]) GetByIndex(index int) *node.Node[T] {
	if index < 0 || index > ll.length {
		return nil
	} else {
		n := ll.head
		for i := 0; i != index; i++ {
			n = n.Next()
		}
		return n
	}
}
