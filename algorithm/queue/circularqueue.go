package queue

import (
	"fmt"
)

type CircularQueue struct {
	queue []interface{}
	cap int
	head int
	tail int
}

func NewCircularQueue(n int) *CircularQueue {
	if n <= 0 {
		return nil
	}
	return &CircularQueue{
		make([]interface{}, n), n, 0, 0,
	}
}

// 循环队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

// 循环队列是否以满
func (q *CircularQueue) IsFull() bool {
	return (q.tail + 1) % q.cap == q.head
}

func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.queue[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
	return true
}

func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.queue[q.head]
	q.head = (q.head + 1) % q.cap
	return v
}

func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}

	result := "head"
	var i = q.head
	for {
		result += fmt.Sprintf("<-%+v", q.queue[i])
		i = (i + 1) % q.cap
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}