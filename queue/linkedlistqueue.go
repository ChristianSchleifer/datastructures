package queue

import "github.com/ChristianSchleifer/datastructures/list"

type linkedListQueue struct {
	list list.List
}

// NewLinkedListQueue returns a linked list based implementation of a queue
func NewLinkedListQueue() Queue {
	return &linkedListQueue{
		list: list.NewSinglyLinkedList(),
	}
}

// Enqueue adds an element to the queue
func (q *linkedListQueue) Enqueue(val interface{}) {
	q.list.Push(val)
}

// Dequeue removes an element from the queue
func (q *linkedListQueue) Dequeue() (interface{}, error) {
	return q.list.Shift()
}
