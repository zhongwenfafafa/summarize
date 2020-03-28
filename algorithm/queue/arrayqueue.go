package queue

import "fmt"

type ArrQueue struct {
	data []interface{}
	cap int
	head int
	tail int
}

func NewArrQueue(cap int) *ArrQueue {
	return &ArrQueue{
		make([]interface{}, cap), cap, 0, 0}
}

func (q *ArrQueue) EnQueue(v interface{}) bool {
	if q.tail == q.cap {
		if q.head == 0 {
			return false
		}
		// 当head不等于0的时候触发搬移操作平均时间复杂度O(1)
		for i := q.head; i < q.tail; i++ {
			q.data[i - q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}

	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	v := q.data[q.head]
	q.head++
	return v
}

func (q *ArrQueue) String() string {
	if q.head == q.tail {
		return "empty queue"
	}

	result := "head"
	for i := q.head; i < q.tail; i++ {
		result += fmt.Sprintf("<-%+v", q.data[i])
	}
	result += "<-tail"
	return result
}