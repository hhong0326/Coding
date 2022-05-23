package main

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) < 1 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		b := []byte(str)
		zeros, ones := 0, 0
		for _, c := range b {
			if c == '1' {
				ones++
			} else {
				zeros++
			}
		}

		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i-zeros][j-ones]+1, dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
