package main

import "math"

func minDistance(word1 string, word2 string) int {

	l1 := len(word1)
	l2 := len(word2)

	// memorization

	//     memo := make([][]int, l1+1)
	//     for i := range memo {
	//         memo[i] = make([]int, l2+1)
	//     }

	//     return l1 + l2 - 2 * lcs(word1, word2, l1, l2, &memo)

	// dp
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			//
			if i == 0 || j == 0 {
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return l1 + l2 - 2*dp[l1][l2]

}

func lcs(s1, s2 string, l, r int, memo *[][]int) int {

	if l == 0 || r == 0 {
		return 0
	}
	if (*memo)[l][r] > 0 {
		return (*memo)[l][r]
	}

	if s1[l-1] == s2[r-1] {
		(*memo)[l][r] = 1 + lcs(s1, s2, l-1, r-1, memo)
	} else {
		(*memo)[l][r] = int(math.Max(float64(lcs(s1, s2, l, r-1, memo)), float64(lcs(s1, s2, l-1, r, memo))))
	}

	return (*memo)[l][r]

}
