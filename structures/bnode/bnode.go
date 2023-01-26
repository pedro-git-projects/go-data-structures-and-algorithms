package bnode

import "github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"

type BNode[T constraints.Ordered] struct {
	value T
	left  *BNode[T]
	right *BNode[T]
}

func New[T constraints.Ordered](value T) *BNode[T] {
	n := &BNode[T]{
		value: value,
		left:  nil,
		right: nil,
	}
	return n
}

func (n BNode[T]) Value() T {
	return n.value
}

func (n BNode[T]) Left() *BNode[T] {
	return n.left
}

func (n *BNode[T]) SetLeft(newNode *BNode[T]) {
	n.left = newNode
}

func (n BNode[T]) Right() *BNode[T] {
	return n.right
}

func (n *BNode[T]) SetRight(newNode *BNode[T]) {
	n.right = newNode
}

func (n *BNode[T]) SetValue(newValue T) {
	n.value = newValue
}
