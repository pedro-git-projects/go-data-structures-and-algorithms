package node

type Node[T any] struct {
	value T
	next  *Node[T]
}

func New[T any](value T) *Node[T] {
	n := &Node[T]{
		value: value,
		next:  nil,
	}
	return n
}

func (n Node[T]) Value() T {
	return n.value
}

func (n Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) SetNext(newNode *Node[T]) {
	n.next = newNode
}
