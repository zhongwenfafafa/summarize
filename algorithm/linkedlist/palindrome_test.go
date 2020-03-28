package linkedlist

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		data string
		t bool
	}{
		{"abccba", true},
		{"abc", false},
		{"aaaaaaa", true},
		{"abcdcba", true},
		{"中文的文中", true},
		{"大号加快立法和", false},
	}

	for i, v := range tests {
		l := NewLinkedList()
		for _, vv := range v.data {
			l.InsertHead(vv)
		}

		result := IsPalindrome(l)
		if v.t !=  result {
			t.Errorf("测试编号%d，预期结果为：%v, 实际结果为：%v",
				i, v.t, result)
		}
	}
}