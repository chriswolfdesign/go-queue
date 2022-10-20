package queue

import "fmt"

type Queue struct {
	Size int
	head *queueNode
	tail *queueNode
}

func (q *Queue) Enqueue(element int) {
	newNode := queueNode{val: element}

	if q.head == nil {
		q.head = &newNode
		q.tail = &newNode
	} else {
		q.tail.next = &newNode
		newNode.prev = q.tail
		q.tail = q.tail.next
	}

	q.Size++
}

func (q *Queue) Peek() (int, error) {
	if q.Size == 0 {
		return 0, fmt.Errorf("cannot peek into empty queue")
	}

	return q.head.val, nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.Size == 0 {
		return 0, fmt.Errorf("cannot dequeue from empty queue")
	}

	val := q.head.val
	q.head = q.head.next

	if q.Size == 1 {
		q.tail = nil
	}

	q.Size--

	return val, nil
}

type queueNode struct {
	val  int
	next *queueNode
	prev *queueNode
}
