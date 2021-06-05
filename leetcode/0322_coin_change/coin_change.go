package coin_change

func coinChange(coins []int, amount int) int {
	// index = amount (0~amount), value = fewest coins
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	result := 0
	if dp[amount] > amount {
		result = -1
	} else {
		result = dp[amount]
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
