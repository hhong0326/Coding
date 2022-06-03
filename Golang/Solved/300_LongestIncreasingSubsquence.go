package main

import "sort"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {

	//     var dp []int

	//     dp = make([]int, len(nums))
	//     for i := range dp {
	//         dp[i] = 1
	//     }

	//     max := 1
	//     for i:=1; i<len(nums); i++ {
	//         for j:=0; j<i; j++ {
	//             if nums[j] < nums[i] {
	//                 dp[i] = Max(dp[i], 1+dp[j])
	//                 max = Max(max, dp[i])
	//             }
	//         }
	//     }

	//     return max

	// other's solution

	var piles [][]int

	for _, num := range nums {

		if len(piles) == 0 {
			piles = append(piles, []int{num})
			continue
		}
		// This is the binary search portion
		// I'm not writing my own binary search, I'll just let Go do it
		// We want to find the left-most pile where the top card is
		// greater than or equal to the new card.
		j := sort.Search(len(piles), func(k int) bool {
			return piles[k][len(piles[k])-1] >= num
		})
		// If we found an appropriate pile, we add to it.
		// Otherwise, we start a new pile.

		if j < len(piles) {
			piles[j] = append(piles[j], num)
		} else {
			piles = append(piles, []int{num})
		}
	}
	// We don't actually care about any of the subsequences or merging the piles.
	// This is simply an early exit from a full-on Patience Sort.
	// The longest subsequence(s) are going to contain cards from every pile
	// so we just need the total number of piles and we have our answer.

	return len(piles)
}
