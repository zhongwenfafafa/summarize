package stack

import "fmt"

// 数组栈
type ArrayList struct {
	data []interface{}
	top int //栈顶指针
}

func NewArrayStack() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0, 32),
		top: -1,
	}
}

// 数据入栈
func (l *ArrayList) Push(v interface{}) {
	if l.top < 0 {
		l.top = 0
	} else {
		l.top += 1
	}

	if l.top > len(l.data) - 1 {
		l.data = append(l.data, v)
	} else {
		l.data[l.top] = v
	}
}

// 数据出栈
func (l *ArrayList) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	v := l.data[l.top]
	l.top -= 1
	return v
}

func (l *ArrayList) Top() interface{} {
	if l.IsEmpty() {
		return nil
	}

	return l.data[l.top]
}

func (l *ArrayList) IsEmpty() bool {
	return l.top < 0
}

func (l *ArrayList) Flush() {
	l.top = -1
}

func (l *ArrayList) Print() {
	if l.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := l.top; i >= 0; i-- {
			fmt.Println(l.data[i])
		}
	}

}