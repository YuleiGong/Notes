package dvide_or_backtrack

import "testing"

func myPow(x float64, n int) float64 {
	var helper func(n int) float64

	helper = func(n int) float64 {
		if n == 0 {
			return 1
		}
		val := helper(n / 2) //注意使用临时变量
		if n%2 == 0 {
			return val * val
		}
		return val * val * x
	}

	if n < 0 {
		return 1 / helper(n)
	}

	return helper(n)
}

func TestMyPow(t *testing.T) {
	t.Log(myPow(0.00001, 2147483647))
}
