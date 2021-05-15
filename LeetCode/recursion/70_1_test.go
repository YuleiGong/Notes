package main

import "testing"

func helper(n int, cache []int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if cache[n] > 0 {
		return cache[n]
	}

	cache[n] = helper(n-1, cache) + helper(n-2, cache)
	return cache[n]

}

func climbStairs(n int) int {
	cache := make([]int, n+1)
	return helper(n, cache)
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(1000))
}
