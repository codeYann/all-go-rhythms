package classic

import "errors"

// Node's definition
type Node struct {
	Key  int
	next *Node
}

// Queue's definition
type Queue struct {
	front  *Node
	rear   *Node
	length int
}

// Functions to use Queue structure

func CreateQueue() *Queue {
	return &Queue{
		front:  nil,
		rear:   nil,
		length: 0,
	}
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) IsEmpty() int {
	if q.length == 0 {
		return 0
	} else {
		return 1
	}
}

func (q *Queue) Frontal() (*Node, error) {
	if q.length == 0 || q.front == nil {
		return nil, errors.New("Queue is empty! It is impossible to return queue front")
	} else {
		return q.front, nil
	}
}

func (q *Queue) Back() (*Node, error) {
	if q.length == 0 || q.front == nil {
		return nil, errors.New("Queue is empty! It is impossible to return queue back")
	} else {
		return q.rear, nil
	}
}

func (q *Queue) Enqueue(value int) {
	node := &Node{Key: value}

	if q.length == 0 || q.front == nil {
		q.front = node
		q.rear = node
	} else {
		q.rear.next = node
		q.rear = node
	}

	q.length += 1
}

func (q *Queue) Dequeue() (*Node, error) {
	if q.length == 0 || q.front == nil {
		return nil, errors.New("Queue is empty!")
	} else {
		temp := q.front
		q.front = q.front.next
		return temp, nil
	}
}
