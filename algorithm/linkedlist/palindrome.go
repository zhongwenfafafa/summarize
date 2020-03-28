package linkedlist

// 开一个栈内存空间存储链表前半段
// 和后半段进行对比
func IsPalindrome(l *LinkedList) bool {
	if l.length <= 1 {
		return true
	}

	center := l.length/2
	// 开辟一个数组存储链表前半段的值
	s := make([]interface{}, 0, center)

	cur := l.head.next
	for i := uint(0); i < center; i++ {
		 s = append(s, cur.value)
		 cur = cur.next
	}

	if l.length%2 != 0 {
		cur = cur.next
		center++
	}

	for cur != nil {
		if cur.value != s[l.length - center - 1] {
			return false
		}
		cur = cur.next
		center++
	}

	return true
}

// 反转链表前半段，和后半段对比
// 最后还原链表
func IsPalindrome2(l *LinkedList) bool {
	return true
}