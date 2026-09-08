package ticketscheduler

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {

	initCap := 3
	queue := NewQueue[int](initCap)

	queue.Enqueue(3)
	queue.Enqueue(7)

	expected := 3
	item, err := queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	expected = 7
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	_, err = queue.Dequeue()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("queue is empty and the error is unrelated : %e", err)
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	newCap := queue.Cap()
	if newCap != 2*initCap {
		t.Errorf("capacity isnt doubled. newCap=%d", newCap)
	}

	expected = 1
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	expected = 2
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	expected = 3
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	expected = 4
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

	queue.Enqueue(245)

	expected = 5
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("an error ocurred : %e", err)
	}
	if item != expected {
		t.Errorf("wront value=%d , it should be %d", item, expected)
	}

}
