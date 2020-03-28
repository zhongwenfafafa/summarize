package stack

import "testing"

func TestArrayList_Push(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Push(4)
	t.Log(s.Pop())
	s.Print()
}

func TestArrayList_Pop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}

func TestArrayList_Top(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
}