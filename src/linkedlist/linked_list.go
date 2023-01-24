package linkedlist

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
		acc += " -> "
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
	if index < 0 || index >= ll.length {
		return nil
	}
	desired := ll.head
	for i := 0; i < index; i++ {
		desired = desired.Next()
	}
	return desired
}

func (ll *LinkedList[T]) Set(index int, value T) bool {
	n := ll.GetByIndex(index)
	if n != nil {
		n.SetValue(value)
		return true
	} else {
		return false
	}
}

func (ll *LinkedList[T]) Remove(index int) bool {
	n := ll.GetByIndex(index)
	if n == nil {
		return false
	}
	if index == 0 {
		b := ll.RemoveFirst()
		return b
	}
	if index == ll.length-1 {
		return ll.RemoveLast()
	}

	prev := ll.GetByIndex(index - 1)
	tmp := prev.Next()
	prev.SetNext(tmp.Next())
	tmp.SetNext(nil)
	ll.length--
	return true
}

func (ll *LinkedList[T]) Insert(index int, value T) bool {
	desired := ll.GetByIndex(index)

	if desired == nil {
		return false
	}

	if index == 0 {
		return ll.Prepend(value)
	}

	if index == ll.length {
		return ll.Append(value)
	}

	n := node.New(value)
	prev := ll.GetByIndex(index - 1)
	n.SetNext(prev.Next())
	prev.SetNext(n)
	ll.length++

	return true
}

func (ll *LinkedList[T]) Reverse() {
	tmp := ll.Head()
	ll.head = ll.tail
	ll.tail = tmp

	after := tmp.Next()
	var before *node.Node[T] = nil

	for i := 0; i < ll.length; i++ {
		after = tmp.Next()
		tmp.SetNext(before)
		before = tmp
		tmp = after
	}
}
