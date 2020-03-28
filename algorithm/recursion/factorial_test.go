package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct{
		n int
		v int
	}{
		{2, 2},
		{3, 6},
		{4, 24},
		{8, 40320},
	}

	for _,test := range tests  {
		v := Factorial(test.n)
		if test.v != v {
			t.Errorf("期望值%d, 实际值%d",
				test.v, v)
		}
	}
}

func TestFac_Factorial(t *testing.T) {
	fac := NewFactorial(10)
	for i := 1; i < 15; i++ {
		fac.Factorial(i)
		fac.Print(i)
	}
}