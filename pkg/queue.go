package queue

type Queue struct {
	Size int
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	val  int
	next *queueNode
	prev *queueNode
}
