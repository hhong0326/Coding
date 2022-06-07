package main

import (
	"fmt"
	"math"
)

const N int = 4

var cols = make([]int, N+1)

func Queens(level int) bool {

	if !condition(level) {
		return false
	} else if level == N {
		for i := 1; i <= N; i++ {

			fmt.Println(i, cols[i])
		}
		return true
	}
	for i := 1; i <= N; i++ {
		cols[level+1] = i
		Queens(level + 1) // 모든 Queens 찾기
		// if Queens(level + 1) { // 하나라도 찾고 끝내기
		// 	return true
		// }
	}

	return false
}

func condition(level int) bool {
	for i := 1; i < level; i++ {
		if cols[level] == cols[i] {
			return false
		} else if level-i == int(math.Abs(float64(cols[level])-float64(cols[i]))) {
			return false
		}
	}
	return true
}

func main() {
	Queens(1)
}
