package queue

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/errs"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/item"
)

type Queue[T any] struct {
	first  *item.Item[T]
	last   *item.Item[T]
	lenght int
}

func New[T any](value T) *Queue[T] {
	i := item.New(value)
	q := &Queue[T]{
		first:  i,
		last:   i,
		lenght: 1,
	}
	return q
}

func (q Queue[T]) First() *item.Item[T] {
	return q.first
}

func (q Queue[T]) Last() *item.Item[T] {
	return q.last
}

func (q Queue[T]) Length() int {
	return q.lenght
}

func (q Queue[T]) String() string {
	str := ""
	tmp := q.first
	for tmp != nil {
		str += fmt.Sprintf("%v", tmp.Value())
		str += " "
		tmp = tmp.Next()
	}
	return str
}

func (q *Queue[T]) Enqueue(value T) {
	i := item.New(value)
	if q.lenght == 0 {
		q.first = i
		q.last = i
	} else {
		q.last.SetNext(i)
		q.last = i
	}
	q.lenght++
}

func (q *Queue[T]) Dequeue() (*item.Item[T], error) {
	if q.lenght == 0 {
		return nil, errs.OpOnZeroLen("deque a", "queue")
	}
	dequeued := q.first
	if q.lenght == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.Next()
	}
	q.lenght--
	return dequeued, nil
}
