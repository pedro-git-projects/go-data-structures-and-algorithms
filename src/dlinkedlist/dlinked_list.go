package dlinkedlist

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/dnode"
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

func (dl *DLinkedList[T]) Append(value T) bool {
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
	return true
}

func (dl *DLinkedList[T]) Prepend(value T) bool {
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
	return true
}

func (dl *DLinkedList[T]) RemoveFirst() bool {
	if dl.length == 0 {
		return false
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
	return true
}

func (dl *DLinkedList[T]) RemoveLast() bool {
	if dl.length == 0 {
		return false
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
	return true
}

func (dl *DLinkedList[T]) Get(index int) *dnode.DNode[T] {
	if index < 0 || index >= dl.length {
		return nil
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

	return desired
}

func (dl *DLinkedList[T]) Set(index int, value T) bool {
	desired := dl.Get(index)
	if desired != nil {
		desired.SetValue(value)
		return true
	}
	return false
}
