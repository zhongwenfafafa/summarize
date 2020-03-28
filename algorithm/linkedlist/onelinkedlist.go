package linkedlist

type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(nil), 0}
}

// 在某个节点后面插入节点
func (l *LinkedList) InsertAfter(n *ListNode, v interface{}) bool {
	if n == nil {
		return false
	}

	newNode := NewListNode(v)
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// 在某个节点前面插入节点
func (l *LinkedList) InsertBefore(n *ListNode, v interface{}) bool {
	if n == nil || n == l.head {
		return false
	}

	newNode := NewListNode(v)
	pre := l.head
	cur := l.head.next

	for cur != nil {
		if cur == n {
			pre.next = newNode
			newNode.next = cur
			l.length++
			return true
		}

		pre = cur
		cur = cur.next
	}

	return false
}

// 根据索引查找指定节点
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if l == nil || l.head == nil || l.head.next == nil {
		return nil
	}

	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}

	return cur
}

// 删除指定节点
func (l *LinkedList) DeleteNode(n *ListNode) bool {
	if n == l.head || n == nil {
		return false
	}

	pre := l.head
	cur := l.head.next
	for cur != nil {
		if cur == n {
			break
		}

		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	pre.next = cur.next
	n.next = nil
	l.length--
	return true
}
