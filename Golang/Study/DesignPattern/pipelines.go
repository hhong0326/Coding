package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // task 가 되는 행동 후 out 채널로 넣으면 된다.
		}
		close(out)
	}()
	return out
}

func main() {
	inputChan := generate(2, 3)
	resultChan := square(inputChan)

	fmt.Println(<-resultChan) // 4
	fmt.Println(<-resultChan) // 9

	for n := range square(square(generate(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
