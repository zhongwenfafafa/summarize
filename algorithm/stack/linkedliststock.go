package stack

import "fmt"

type ListNode struct {
	next *ListNode
	value interface{}
}

func NewLinkedListStack() *LinkedList {
	return &LinkedList{nil}
}

type LinkedList struct {
	topNode *ListNode
}

func (l *LinkedList) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}

	v := l.topNode.value
	l.topNode = l.topNode.next
	return v
}

func (l *LinkedList) Push(v interface{}) {
	newNode := &ListNode{l.topNode, v}
	l.topNode = newNode
}

func (l *LinkedList) IsEmpty() bool {
	return l.topNode == nil
}

func (l *LinkedList) Flush () {
	l.topNode = nil
}

func (l *LinkedList) Top() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.topNode.value
}

func (l *LinkedList) Print() {
	if l.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := l.topNode
		for cur != nil {
			fmt.Println(cur.value)
			cur = cur.next
		}
	}
}