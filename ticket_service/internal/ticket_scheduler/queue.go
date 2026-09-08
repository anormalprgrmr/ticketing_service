package ticketscheduler

import (
	"errors"
	"ticket_service/internal/models"
)

var ErrEmpty = errors.New("queue is empty")

type Queue struct {
	items []models.Ticket
	head  int
	size  int
}

func NewQueue(capacity int) *Queue {
	if capacity < 1 {
		capacity = 16
	}

	return &Queue{
		items: make([]models.Ticket, capacity),
	}
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Cap() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Enqueue(ticket models.Ticket) {
	if q.size == len(q.items) {
		q.grow()
	}

	tail := (q.head + q.size) % len(q.items)
	q.items[tail] = ticket
	q.size++
}

func (q *Queue) Dequeue() (models.Ticket, error) {
	if q.size == 0 {
		return models.Ticket{}, ErrEmpty
	}

	ticket := q.items[q.head]

	q.items[q.head] = models.Ticket{}

	q.head = (q.head + 1) % len(q.items)
	q.size--

	if q.size == 0 {
		q.head = 0
	}

	return ticket, nil
}

func (q *Queue) Peek() (models.Ticket, error) {
	if q.size == 0 {
		return models.Ticket{}, ErrEmpty
	}

	return q.items[q.head], nil
}

func (q *Queue) grow() {
	newCapacity := len(q.items) * 2
	if newCapacity == 0 {
		newCapacity = 16
	}

	newItems := make([]models.Ticket, newCapacity)

	for i := 0; i < q.size; i++ {
		index := (q.head + i) % len(q.items)
		newItems[i] = q.items[index]
	}

	q.items = newItems
	q.head = 0
}
