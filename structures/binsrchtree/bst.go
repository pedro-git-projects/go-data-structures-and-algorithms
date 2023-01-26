package binsrchtree

import (
	"reflect"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/bnode"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/errs"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/utils/constraints"
)

type BST[T constraints.Ordered] struct {
	root *bnode.BNode[T]
}

func New[T constraints.Ordered](value T) *BST[T] {
	r := bnode.New(value)
	bst := &BST[T]{
		root: r,
	}
	return bst
}

func (bst BST[T]) Root() *bnode.BNode[T] {
	return bst.root
}

func (bst BST[T]) String() string {
	str := ""
	return str
}

func (bst *BST[T]) Insert(value T) error {
	n := bnode.New(value)
	if bst.root == nil {
		bst.root = n
		return nil
	}

	tmp := bst.root
	for {
		if reflect.DeepEqual(n.Value(), tmp.Value()) {
			return errs.Duplicated
		}
		if n.Value() < tmp.Value() {
			if tmp.Left() == nil {
				tmp.SetLeft(n)
				return nil
			}
			tmp = tmp.Left()
		} else {
			if tmp.Right() == nil {
				tmp.SetRight(n)
				return nil
			}
			tmp = tmp.Right()
		}
	}
}

func (bst BST[T]) Contains(value T) bool {
	if bst.root == nil {
		return false
	}

	tmp := bst.root
	for tmp != nil {
		if value < tmp.Value() {
			tmp = tmp.Left()
		} else if value > tmp.Value() {
			tmp = tmp.Right()
		} else {
			return true
		}
	}

	return false
}
