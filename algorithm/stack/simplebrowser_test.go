package stack

import "testing"

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.PushBack("www.qq.com")
	b.PushBack("www.baidu.com")
	b.PushBack("www.sina.com")

	b.Back()
	b.Forward()

	b.Back()
	b.Back()
	b.Back()

	b.Open("www.taobao.com")
	b.Forward()
	b.Back()
}