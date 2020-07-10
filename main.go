package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====", maxProfit([]int{1, 2, -10, 0, 2}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = 0          // 不持股
	dp[0][1] = -prices[0] // 持股
	dp[0][2] = 0          // 冷冻期
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
