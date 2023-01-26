package linkedlist

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/errs"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/node"
)

// TODO remove return values from functions that can't return errors

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

func (ll *LinkedList[T]) Append(value T) {
	n := node.New(value)
	if ll.length == 0 {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.SetNext(n)
		ll.tail = n
	}
	ll.length++
}

func (ll *LinkedList[T]) RemoveLast() error {
	if ll.length == 0 {
		return errs.OpOnZeroLen("append", "list")
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

	return nil
}

func (ll *LinkedList[T]) Prepend(value T) {
	n := node.New(value)
	if ll.head != nil {
		n.SetNext(ll.head)
		ll.head = n
	} else {
		ll.head = n
		ll.tail = n
	}
	ll.length++
}

func (ll *LinkedList[T]) RemoveFirst() error {
	if ll.length == 0 {
		return errs.OpOnZeroLen("remove first in", "list")
	}

	desired := ll.head
	ll.head = ll.head.Next()
	desired.SetNext(nil)
	ll.length--

	if ll.length == 0 {
		ll.tail = nil
	}
	return nil
}

func (ll *LinkedList[T]) GetByIndex(index int) (*node.Node[T], error) {
	if index < 0 || index >= ll.length {
		return nil, errs.OutOfBounds
	}
	desired := ll.head
	for i := 0; i < index; i++ {
		desired = desired.Next()
	}
	return desired, nil
}

func (ll *LinkedList[T]) Set(index int, value T) error {
	n, err := ll.GetByIndex(index)
	if n != nil {
		n.SetValue(value)
		return nil
	} else {
		return err
	}
}

func (ll *LinkedList[T]) Remove(index int) error {
	_, err := ll.GetByIndex(index)
	if err != nil {
		return err
	}
	if index == 0 {
		b := ll.RemoveFirst()
		return b
	}
	if index == ll.length-1 {
		return ll.RemoveLast()
	}

	prev, _ := ll.GetByIndex(index - 1)
	tmp := prev.Next()

	prev.SetNext(tmp.Next())
	tmp.SetNext(nil)

	ll.length--
	return nil
}

func (ll *LinkedList[T]) Insert(index int, value T) error {
	_, err := ll.GetByIndex(index)

	if err != nil {
		return err
	}

	if index == 0 {
		ll.Prepend(value)
		return nil
	}

	if index == ll.length {
		ll.Append(value)
		return nil
	}

	n := node.New(value)
	prev, _ := ll.GetByIndex(index - 1)
	n.SetNext(prev.Next())
	prev.SetNext(n)
	ll.length++

	return nil
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
