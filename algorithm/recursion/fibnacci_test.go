package recursion

import "testing"

func TestFib_Fibonacci(t *testing.T) {
	fib := NewFibonacci(10)
	for i := 1; i < 15; i++ {
		fib.Fibonacci(i)
		fib.Print(i)
	}
}