package main

import "testing"

func climbStairs(n int) int {
	var cache = make([]int, n+1)
	var helper func(n int) int

	helper = func(n int) int {
		if n == 1 || n == 2 {
			return n
		}
		if cache[n] != 0 {
			return cache[n]
		}
		cache[n] = helper(n-1) + helper(n-2)
		return cache[n]
	}

	return helper(n)
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(20))
}
