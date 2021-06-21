package greedy

import "testing"

func maxProfit(prices []int) int {
	var result int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result

}

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit(prices))
}
