package recursion

// 递归实现阶乘
func Factorial(n int) int {
	if n <= 1 {
		return n
	}
	v := n * Factorial(n - 1)
	return v
}

type Fac struct {
	val map[int]int
}

func NewFactorial(n int) *Fac {
	return &Fac{
		make(map[int]int, n),
	}
}

func (f *Fac) Factorial(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}

	if n <= 1 {
		f.val[n] = 1
		return 1
	}

	v := n * f.Factorial(n - 1)
	f.val[n] = v
	return v
}

func (f *Fac) Print(n int) {
	println(f.val[n])
}