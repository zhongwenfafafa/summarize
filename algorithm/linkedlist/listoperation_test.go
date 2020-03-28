package linkedlist

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct{
		before string
		after string
		t bool
	}{
		{"abccba", "abccba",true},
		{"abc", "cba",false},
		{"aaaaaaa", "aaaaaaa", true},
		{"abcdcba", "abcdcba",true},
		{"测考测感觉哦加哦中文的文中", "中文的文中哦加哦觉感测考测",true},
		{"大号加快立法和", "和法立快加号大",false},
	}

	for i, v := range tests {
		l := NewLinkedList()
		for _, vv := range []rune(v.before) {
			l.InsertTail(string(vv))
		}

		l.Reverse()
		result := l.Print("")
		if v.after !=  result {
			t.Errorf("测试编号%d，预期结果为：%s, 实际结果为：%s",
				i, v.after, result)
		}
	}
}

func TestMergeSortList(t *testing.T) {
	tests := []struct{
		l1 []int
		l2 []int
		result string
	} {
		{[]int{1,2,3}, []int{4,5,6}, "1,2,3,4,5,6"},
		{[]int{1,8,15}, []int{4,13,14}, "1,4,8,13,14,15"},
		{[]int{1,6,7}, []int{4,6,8}, "1,4,6,6,7,8"},
	}

	for i, v := range tests {
		l1 := NewLinkedList()
		l2 := NewLinkedList()
		for _,vv := range v.l1 {
			l1.InsertTail(vv)
		}

		for _,vv := range v.l2  {
			l2.InsertTail(vv)
		}

		megl := MergeSortList(l1, l2)
		result := megl.Print(",")
		if result != v.result {
			t.Errorf("测试编号%d，预期结果为：%s, 实际结果为：%s",
				i, v.result, result)
		}
	}
}

func TestDeleteBottomN(t *testing.T) {
	tests := []struct{
		l1 []int
		index int
		result string
	} {
		{[]int{1,2,3,4,5,6}, 1,"1,2,3,4,5"},
		{[]int{1,8,15,4,13,14}, 3,"1,8,15,13,14"},
		{[]int{1,6,7,4,6,8}, 5,"1,7,4,6,8"},
		{[]int{1,6,7,4,6,8}, 8,"1,6,7,4,6,8"},
		{[]int{1,6,7,4,6,8,8,9}, 0,"1,6,7,4,6,8,8,9"},
	}

	for i, v := range tests {
		l1 := NewLinkedList()
		for _,vv := range v.l1 {
			l1.InsertTail(vv)
		}

		result := l1.DeleteBottomN(v.index).Print(",")
		if result != v.result {
			t.Errorf("测试编号%d，预期结果为：%s, 实际结果为：%s",
				i, v.result, result)
		}
	}
}

func TestFindMiddleNode(t *testing.T) {
	tests := []struct{
		l1 []int
		result int
	} {
		{[]int{1,2,3,4,5,6}, 3},
		{[]int{1,8,15,4,13,14}, 15},
		{[]int{1,6,7,4,6,8}, 7},
		{[]int{1,6,7,4,6,8}, 7},
		{[]int{1,6,7,4,6,8,8,9}, 4},
	}

	for i, v := range tests {
		l1 := NewLinkedList()
		for _,vv := range v.l1 {
			l1.InsertTail(vv)
		}

		result := l1.FindMiddleNode()
		if result.GetValue() != v.result {
			t.Errorf("测试编号%d，预期结果为：%d, 实际结果为：%d",
				i, v.result, result.GetValue())
		}
	}
}