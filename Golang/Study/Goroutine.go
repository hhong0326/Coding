package main

import (
	"fmt"
	"time"
)

func countingStar(d chan string, num string) {

	defer close(d)
	for i := 0; i < 10; i++ {
		d <- num
		time.Sleep(1 * time.Second)
	}
}

func sum(arr []int, c chan int) {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {

	// c := make(chan int)
	// d := make(chan string)
	// go countingStar(d, "1")
	// // go countingStar(d, "2")
	// // go countingStar(d, "3")
	// // for cc := range c {
	// // 	fmt.Println(cc)
	// // }
	// s := time.Now()
	// for dd := range d {
	// 	fmt.Println(dd)
	// }

	// fmt.Println(time.Since(s))

	// 합을 구하는 방법을 고루틴으로 분산하기
	var arr [1000000]int
	for i := 0; i < 1000000; i++ {
		arr[i] = i
	}
	s1 := time.Now()
	cc := make(chan int)

	go sum(arr[:len(arr)/2], cc)
	go sum(arr[len(arr)/2:], cc)
	aSum, bSum := <-cc, <-cc
	fmt.Println(aSum + bSum)
	fmt.Println(time.Since(s1))

	s2 := time.Now()

	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(time.Since(s2))

	//약 10만개부터 고루틴을 사용하여 분산한 결과가 빨라짐
}
