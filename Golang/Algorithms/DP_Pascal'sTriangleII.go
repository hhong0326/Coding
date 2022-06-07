package main

func getRow(rowIndex int) []int {

	dp := make([][]int, rowIndex+1)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = 1

	for i := 1; i < len(dp); i++ {
		dp[i][0] = 1
		dp[i][len(dp[i])-1] = 1
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]

		}

	}

	return dp[rowIndex]
}
