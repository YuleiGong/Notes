package dvide_or_backtrack

import "testing"

func myPow(x float64, n int) float64 {
	var helper func(n int) float64

	helper = func(n int) float64 {
		if n == 0 {
			return 1
		}
		val := helper(n / 2)
		if n%2 == 0 {
			return val * val
		}
		return val * val * x //奇数
	}
	if n >= 0 {
		return helper(n)
	}
	return 1 / helper(-n)

}

func TestMyPow(t *testing.T) {
	t.Log(myPow(2, 3))
}
