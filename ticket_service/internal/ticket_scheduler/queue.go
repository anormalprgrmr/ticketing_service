package ticketscheduler

import (
	"errors"
)

var ErrEmpty = errors.New("queue is empty")

type Queue[T any] struct {
	items []T
	head  int
	size  int
}

func NewQueue[T any](capacity int) *Queue[T] {
	if capacity < 1 {
		capacity = 16
	}

	return &Queue[T]{
		items: make([]T, capacity),
	}
}

func (q *Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) Cap() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Enqueue(item T) {
	if q.size == len(q.items) {
		q.grow()
	}

	tail := (q.head + q.size) % len(q.items)
	q.items[tail] = item
	q.size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.size == 0 {
		return zero, ErrEmpty
	}

	item := q.items[q.head]

	q.items[q.head] = zero

	q.head = (q.head + 1) % len(q.items)
	q.size--

	if q.size == 0 {
		q.head = 0
	}

	return item, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.size == 0 {
		var zero T
		return zero, ErrEmpty
	}

	return q.items[q.head], nil
}

func (q *Queue[T]) grow() {
	newCapacity := len(q.items) * 2
	if newCapacity == 0 {
		newCapacity = 16
	}

	newItems := make([]T, newCapacity)

	for i := 0; i < q.size; i++ {
		index := (q.head + i) % len(q.items)
		newItems[i] = q.items[index]
	}

	q.items = newItems
	q.head = 0
}
