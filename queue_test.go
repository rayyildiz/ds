package ds

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	if q.Length() != 0 {
		t.Error("Error creating new queue")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue("One")

	if q.Length() != 1 {
		t.Error("Error at enqueue")
	}
}

func TestEnqueue2(t *testing.T) {
	q := NewQueue()

	q.Enqueue("One")
	q.Enqueue("Two")

	if q.Length() != 2 {
		t.Error("Error at enqueue")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue("One")
	q.Enqueue("Two")
	q.Enqueue("Three")

	elem, err := q.Dequeue()
	if err != nil {
		t.Error("Error at dequeue 'One'")
	}

	if elem != "One" {
		t.Error("Error at dequeue 'One'")
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Error("Error at dequeue 'Two'")
	}

	if elem != "Two" {
		t.Error("Error at dequeue 'Two'")
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Error("Error at dequeue 'Three'")
	}

	if elem != "Three" {
		t.Error("Error at dequeue 'Three'")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Error("Error at dequeue for empty")
	}
}
