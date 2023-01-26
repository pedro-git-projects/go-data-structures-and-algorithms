package hnode

type HNode[T any] struct {
	key   any
	value T
	next  *HNode[T]
}

func New[T any](key any, value T) *HNode[T] {
	n := &HNode[T]{
		key:   key,
		value: value,
		next:  nil,
	}
	return n
}

func (n HNode[T]) Value() T {
	return n.value
}

func (n HNode[T]) Key() any {
	return n.key
}

func (n HNode[T]) Next() *HNode[T] {
	return n.next
}

func (n *HNode[T]) SetNext(newNode *HNode[T]) {
	n.next = newNode
}

func (n *HNode[T]) SetKey(newKey T) {
	n.key = newKey
}

func (n *HNode[T]) SetValue(newValue T) {
	n.value = newValue
}
