package main

import (
	"math"
)

var (
	m      map[string]string
	profit map[string]int
)

func solution(enroll []string, referral []string, seller []string, amount []int) []int {

	answer := []int{}

	m = make(map[string]string, 0)
	profit = make(map[string]int, 0)

	for i := 0; i < len(enroll); i++ {
		ref := referral[i]
		enr := enroll[i]

		m[enr] = ref
	}

	for i := 0; i < len(seller); i++ {
		sel := seller[i]
		v := m[sel]

		p := profit[sel]
		price := amount[i] * 100

		profit[sel] = p + price - int(math.Floor(float64(price)*0.1))

		dfs(v, int(math.Floor(float64(price)*0.1)))
	}

	for i := 0; i < len(enroll); i++ {
		v := profit[enroll[i]]

		answer = append(answer, v)
	}

	return answer

}

func dfs(parent string, remain int) {
	if parent == "-" || remain < 1 {
		return
	}

	price := profit[parent]
	v := m[parent]

	profit[parent] = price + remain - int(math.Floor(float64(remain)*0.1))

	dfs(v, int(math.Floor(float64(remain)*0.1)))

}
