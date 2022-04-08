package main

import (
	"math"
)

func main() {
	solution(10, 10, []int{100}, []int{100}, []int{7}, []int{10})
}

func solution(a int, b int, g []int, s []int, w []int, t []int) int64 {

	// a, b는 새로운 도시에 필요한 광물
	// 기존 도시가 가지고 있는 g와 s, 트럭 하중 w, 새로운 도시 편도 시간 t
	// 가장 빠른 시간

	answer := -1
	// 최악의 경우: 최대의 광물 + 최대의 시간 + a
	// 10의 9 * 2 * 10의 5 * 2
	l, r := 0, int(1e9*2*1e5*2) // math.MaxInt 9223372036854775807
	answer = r

	for l <= r {
		mid := (l + r) / 2
		if search(a, b, g, s, w, t, mid) {
			answer = int(math.Min(float64(answer), float64(mid)))
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int64(answer)
}

func search(a, b int, g, s []int, w, t []int, mid int) bool {
	totalGold := 0
	totalSilver := 0
	totalMineral := 0

	for i := 0; i < len(g); i++ {
		t1 := t[i]
		round := t1 * 2
		move := mid / round
		if mid%round >= t1 {
			move++
		}
		max := w[i] * move

		totalGold += int(math.Min(float64(g[i]), float64(max)))
		totalSilver += int(math.Min(float64(s[i]), float64(max)))
		totalMineral += int(math.Min(float64(g[i]+s[i]), float64(max)))
	}

	if totalGold >= a && totalSilver >= b && totalMineral >= a+b {
		return true
	}

	return false
}

// 참고: https://yabmoons.tistory.com/714
