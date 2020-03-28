package stack

import "fmt"

// 模拟游览器前进后退简单实现
type Browser struct {
	forwardStack Stack
	backStack Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack: NewLinkedListStack(),
	}
}

// 是否能前进
func (b *Browser) CanForward() bool {
	return !b.forwardStack.IsEmpty()
}

// 是否能后退
func (b *Browser) CanBack() bool {
	return !b.backStack.IsEmpty()
}

// 打开新的页面
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forwardStack.Flush()
}

func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

func (b *Browser) Forward() {
	if b.CanForward() {
		top := b.forwardStack.Pop()
		b.backStack.Push(top)
		fmt.Printf("forward to %+v\n", top)
		return
	}
	fmt.Println("can't forward")
}

func (b *Browser) Back() {
	if b.CanBack() {
		top := b.backStack.Pop()
		b.forwardStack.Push(top)
		fmt.Printf("back to %+v\n", top)
		return
	}
	fmt.Println("can't back")
}