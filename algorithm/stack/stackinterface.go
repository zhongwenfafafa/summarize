package stack

type Stack interface {
	Pop() interface{}
	Push(v interface{})
	Top() interface{}
	IsEmpty() bool
	Flush()
}