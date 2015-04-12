package ds

import (
	"container/list"
	"errors"
)

// ErrQueue is for error if queue is empty.
var ErrQueue = errors.New("queue is empty")

// Queue Data Container
type Queue struct {
	container *list.List
}

// NewQueue : Create a new queue
func NewQueue() *Queue {
	q := Queue{}
	q.container = list.New()

	return &q
}

// Enqueue : Insert an item to queue
func (q *Queue) Enqueue(value interface{}) {
	q.container.PushFront(value)
}

// Dequeue : Remove an item from queue (FIFO)
func (q *Queue) Dequeue() (interface{}, error) {
	if q.Length() < 1 {
		return nil, ErrQueue
	}

	elem := q.container.Back()
	val := q.container.Remove(elem)

	return val, nil
}

// Length of the queue
func (q *Queue) Length() int {
	return q.container.Len()
}
