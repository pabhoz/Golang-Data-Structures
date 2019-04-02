package queue

import (
	"fmt"

	dll "../DLL"
)

type (
	//Queue Basic queue struct
	Queue struct {
		front, rear *dll.Element
		elements    *dll.List
		max         int
	}
)

// New Creates a new Queue
func New(max int) *Queue {
	list := dll.New()
	return &Queue{nil, nil, list, max}
}

// Front returns front queue element
func (q *Queue) Front() *dll.Element {
	return q.front
}

// Rear returns rear queue element
func (q *Queue) Rear() *dll.Element {
	return q.rear
}

// IsEmpty returns true or false
func (q *Queue) IsEmpty() bool {
	return q.elements.Len() == 0
}

// IsFull returns true or false
func (q *Queue) IsFull() bool {
	return q.elements.Len() == q.max
}

// Enqueue add a new element to the quee
func (q *Queue) Enqueue(value interface{}) *dll.Element {
	el := q.elements.PushFront(value)
	q.reorder()
	return el
}

// Dequeue add a new element to the quee
func (q *Queue) Dequeue() interface{} {
	toRemove := q.elements.Back()
	q.elements.Remove(toRemove)
	q.reorder()
	return toRemove.Value
}

func (q *Queue) reorder() {
	q.front = q.elements.Back()
	q.rear = q.elements.Front()
}

//PrintMe  prints the queue
func (q *Queue) PrintMe() {
	fmt.Println("Printing Queue (BFS):")
	for e := q.elements.Back(); e != nil; e = e.Prev() {
		fmt.Println(e)
	}
}
