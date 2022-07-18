package main

import (
	"fmt"
	"time"
)

func main() {

	s1 := time.Now()
	res := []int{}
	x, y := 0, 1
	for i := 0; i <= 10; i++ {
		res = append(res, x)
		x, y = y, x+y

	}
	fmt.Println(time.Since(s1))
	fmt.Println(res)

	s2 := time.Now()
	res2 := make([]int, 11)
	res2[0], res2[1] = 0, 1

	for i := 2; i <= 10; i++ {
		res2[i] = res2[i-1] + res2[i-2]

	}
	fmt.Println(time.Since(s2))
	fmt.Println(res2)
}
