package item

type Item[T any] struct {
	value T
	next  *Item[T]
}

func New[T any](value T) *Item[T] {
	n := &Item[T]{
		value: value,
		next:  nil,
	}
	return n
}

func (i Item[T]) Value() T {
	return i.value
}

func (i Item[T]) Next() *Item[T] {
	return i.next
}

func (i *Item[T]) SetNext(newNode *Item[T]) {
	i.next = newNode
}

func (i *Item[T]) SetValue(newValue T) {
	i.value = newValue
}
