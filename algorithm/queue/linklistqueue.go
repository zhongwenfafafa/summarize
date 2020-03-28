package queue

import "fmt"

type LinkedListQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

type ListNode struct {
	next *ListNode
	value interface{}
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{nil, v}
	if nil == q.tail {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}

	v := q.head.value
	q.head = q.head.next
	q.length--
	return v
}

func (q *LinkedListQueue) String() string {
	if q.head == nil {
		return "empty queue"
	}

	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.value)

	}
	result += "<-tail"

	return result
}