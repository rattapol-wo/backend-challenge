package challenge

func SumValue(value [][]int) int {
	n := len(value)
	if n == 0 {
		return 0
	}

	// Use a copy of the last row to store the maximum path sum
	dp := make([]int, n)
	copy(dp, value[n-1])

	// Start from the second-to-last row and move upwards
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// Calculate the maximum value between the left and right paths
			dp[j] = value[i][j] + max(dp[j], dp[j+1])
		}
	}

	// The maximum path sum will be at the first position
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}