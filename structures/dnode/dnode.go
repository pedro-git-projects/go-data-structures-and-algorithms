package dnode

type DNode[T any] struct {
	value T
	next  *DNode[T]
	prev  *DNode[T]
}

func New[T any](value T) *DNode[T] {
	n := &DNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}
	return n
}

func (n DNode[T]) Value() T {
	return n.value
}

func (n DNode[T]) Next() *DNode[T] {
	return n.next
}

func (n DNode[T]) Prev() *DNode[T] {
	return n.prev
}

func (n *DNode[T]) SetNext(newNode *DNode[T]) {
	n.next = newNode
}

func (n *DNode[T]) SetPrev(newNode *DNode[T]) {
	n.prev = newNode
}

func (n *DNode[T]) SetValue(newValue T) {
	n.value = newValue
}
