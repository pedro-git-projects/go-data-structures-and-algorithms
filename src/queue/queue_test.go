package queue_test

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/queue"
)

func TestEnqueue(t *testing.T) {
	q := queue.New(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expect := 3
	got := q.Length()

	if expect != got {
		t.Errorf("Expected %d but got %d", expect, got)
	}

	got = q.Last().Value()
	expect = 3

	if expect != got {
		t.Errorf("Expected %d but got %d", expect, got)
	}

	got = q.First().Value()
	expect = 1

	if expect != got {
		t.Errorf("Expected %d but got %d", expect, got)
	}
}

func TestDequeue(t *testing.T) {
	// len 1
	q := queue.New(1)
	q.Dequeue()

	got := q.First()
	if got != nil {
		t.Errorf("Did not expect value but got %v", got)
	}

	expect := 0
	length := q.Length()
	if expect != length {
		t.Errorf("Expected %d but got %d", expect, length)
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected an error but got none")
	}

	// len > 1
	qq := queue.New(1)
	qq.Enqueue(2)
	qq.Enqueue(3)

	result, err := qq.Dequeue()

	if err != nil {
		t.Error(err)
	}

	expect = 1
	if expect != result.Value() {
		t.Errorf("Expected %d but got %d", expect, length)
	}

	expect = 2
	length = qq.Length()
	if expect != length {
		t.Errorf("Expected %d but got %d", expect, length)
	}
}
