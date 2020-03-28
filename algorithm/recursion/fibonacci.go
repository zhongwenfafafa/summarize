package recursion

type Fib struct {
	val map[int]int
}

func NewFibonacci(n int) *Fib {
	return &Fib{make(map[int]int, n)}
}

func (f *Fib) Fibonacci(n int) int {
	if _, ok := f.val[n]; ok {
		return f.val[n]
	}

	if n <= 1 || n == 2 {
		f.val[n] = 1
		return 1
	}

	v := f.Fibonacci(n - 1) + f.Fibonacci(n - 2)
	f.val[n] = v
	return v
}

func (f *Fib) Print(n int) {
	println(f.val[n])
}