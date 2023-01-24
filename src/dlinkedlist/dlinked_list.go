package dlinkedlist

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/dnode"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/errs"
)

type DLinkedList[T any] struct {
	head   *dnode.DNode[T]
	tail   *dnode.DNode[T]
	length int
}

func New[T any](value T) *DLinkedList[T] {
	n := dnode.New(value)
	dl := &DLinkedList[T]{
		head:   n,
		tail:   n,
		length: 1,
	}
	return dl
}

func (dl DLinkedList[T]) Head() *dnode.DNode[T] {
	return dl.head
}

func (dl DLinkedList[T]) Tail() *dnode.DNode[T] {
	return dl.tail
}

func (dl DLinkedList[T]) Length() int {
	return dl.length
}

func (dl DLinkedList[T]) String() string {
	temp := dl.head
	acc := ""
	for temp != nil {
		acc += fmt.Sprintf("%v", temp.Value())
		acc += " "
		temp = temp.Next()
	}
	return acc
}

func (dl *DLinkedList[T]) Append(value T) error {
	n := dnode.New(value)
	if dl.head == nil {
		dl.head = n
		dl.tail = n
	} else {
		dl.tail.SetNext(n)
		n.SetPrev(dl.Tail())
		dl.tail = n
	}
	dl.length++
	return nil
}

func (dl *DLinkedList[T]) Prepend(value T) error {
	n := dnode.New(value)
	if dl.length == 0 {
		dl.head = n
		dl.tail = n
	} else {
		n.SetNext(dl.head)
		dl.head.SetPrev(n)
		dl.head = n
	}
	dl.length++
	return nil
}

func (dl *DLinkedList[T]) RemoveFirst() error {
	if dl.length == 0 {
		return errs.OpOnZeroLen("remove first in", "list")
	}

	desired := dl.head

	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = dl.head.Next()
		dl.head.SetPrev(nil)
		desired.SetNext(nil)
	}
	dl.length--
	return nil
}

func (dl *DLinkedList[T]) RemoveLast() error {
	if dl.length == 0 {
		return errs.OpOnZeroLen("remove last in", "list")
	}

	desired := dl.tail

	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail = desired.Prev()
		dl.tail.SetNext(nil)
		desired.SetPrev(nil)
	}
	dl.length--
	return nil
}

func (dl *DLinkedList[T]) Get(index int) (*dnode.DNode[T], error) {
	if index < 0 || index >= dl.length {
		return nil, errs.OutOfBounds
	}

	desired := dl.head
	if dl.length < index/2 {
		for i := 0; i < index; i++ {
			desired = desired.Next()
		}
	} else {
		desired = dl.tail
		for i := dl.length - 1; i > index; i-- {
			desired = desired.Prev()
		}
	}

	return desired, nil
}

func (dl *DLinkedList[T]) Set(index int, value T) error {
	desired, err := dl.Get(index)
	if err == nil {
		desired.SetValue(value)
		return nil
	}
	return err
}

func (dl *DLinkedList[T]) Insert(index int, value T) error {
	_, err := dl.Get(index)
	if err != nil {
		return err
	}

	if index == 0 {
		return dl.Prepend(value)
	}

	if index == dl.length {
		return dl.Append(value)
	}

	n := dnode.New(value)
	before, _ := dl.Get(index - 1)
	after := before.Next()

	n.SetPrev(before)
	n.SetNext(after)
	before.SetNext(n)
	after.SetPrev(n)
	dl.length++
	return nil
}

func (dl *DLinkedList[T]) Remove(index int) error {
	desired, err := dl.Get(index)
	if err != nil {
		return err
	}

	if index == 0 {
		return dl.RemoveFirst()
	}

	if index == dl.length-1 {
		return dl.RemoveLast()
	}

	desired.Next().SetPrev(desired.Prev())
	desired.Prev().SetNext(desired.Next())
	desired.SetNext(nil)
	desired.SetPrev(nil)
	dl.length--

	return nil
}
