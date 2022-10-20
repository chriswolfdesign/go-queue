package queue

import "fmt"

type Queue struct {
	size int
	head *queueNode
	tail *queueNode
}

func (q *Queue) Size() int {
	return q.size
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

	q.size++
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("cannot peek into empty queue")
	}

	return q.head.val, nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("cannot dequeue from empty queue")
	}

	val := q.head.val
	q.head = q.head.next

	if q.size == 1 {
		q.tail = nil
	}

	q.size--

	return val, nil
}

type queueNode struct {
	val  int
	next *queueNode
	prev *queueNode
}
