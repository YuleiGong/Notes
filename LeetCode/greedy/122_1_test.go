package greedy

import "testing"

func maxProfit(prices []int) int {
	var result int

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			result += profit
		}
	}

	return result

}

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit(prices))
}
