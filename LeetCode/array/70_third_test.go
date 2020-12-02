package array

import "testing"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	var f3 int
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f2, f1 = f3, f2
	}

	return f3
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(4))
}
