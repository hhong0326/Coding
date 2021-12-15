package coding

import "fmt"

func maxProfit(prices []int) int {

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	fmt.Println("2021.11.11")

	return profit
}
