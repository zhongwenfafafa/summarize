package linkedlist_1

import "fmt"

type LinkNode struct {
	next *LinkNode
	data interface{}
}

type LinkedList struct {
	head *LinkNode // 链表的头节点
}

func (n *LinkNode) GetValue() interface{} {
	return n.data
}

// 创建一个新节点
func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		NewLinkNode(0),
	}
}


// 向链表开头插入节点
func (l *LinkedList) InsertHead(v interface{}) bool {
	cur := l.head
	return l.InsertAfter(cur, v)
}

// 向链表最后插入节点
func (l *LinkedList) InsertTail(v interface{}) bool {
	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

// 在某个节点后面插入节点
func (l *LinkedList) InsertAfter(n *LinkNode, v interface{}) bool {
	if n == nil {
		return false
	}

	newNode := NewLinkNode(v)
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	return true
}

// 打印链表
func (l *LinkedList) Print(option string) string {
	cur := l.head.next
	format := ""

	for cur != nil {
		format += fmt.Sprintf("%v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += option
		}
	}
	return format
	//fmt.Println(format)
}

// 单链表反转
func (l *LinkedList) Reverse() {
	// 空链表或链表只有一个元素直接返回
	if l.head.next == nil || l.head == nil || l.head.next.next == nil {
		return
	}

	// 执行反转
	cur := l.head.next
	var pre *LinkNode = nil
	for cur != nil {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}

	l.head.next = pre
}

// 检测链表中是否有环
func (l *LinkedList) HasCycle() bool {
	if nil != l.head && l.head.next != nil {
		slow := l.head.next
		fast := l.head.next.next
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}

	return false
}

// 两个有序单链表合并
func MergeSortList(l1, l2 *LinkedList) *LinkedList {
	if l1.head == nil || l1.head.next == nil || l1 == nil {
		return l2
	}

	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	// 初始化一个新链表
	l := &LinkedList{head: &LinkNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for curl1 != nil && curl2 != nil {
		if curl1.GetValue().(int) > curl2.GetValue().(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}

		cur = cur.next
	}

	if curl1 != nil {
		cur.next = curl1
	} else if curl2 != nil {
		cur.next = curl2
	}

	return l
}

// 删除链表的倒数第N个节点
func (l *LinkedList) DeleteBottomN(n int) *LinkedList{
	// 排除链表为nil或为空的情况
	if l == nil || l.head == nil || l.head.next == nil || n <= 0 {
		return l
	}

	fast := l.head.next
	for i := 1; i < n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return l
	}

	slow := l.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	return l
}

// 获取链表的中间节点
func (l *LinkedList) FindMiddleNode() *LinkNode {
	if nil == l || l.head == nil || l.head.next == nil {
		return nil
	}

	if nil == l.head.next.next {
		return l.head.next
	}

	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}