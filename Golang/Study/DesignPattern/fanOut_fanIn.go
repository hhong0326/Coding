package main

import (
	"fmt"
	"sync"
)

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

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	multiplex := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	inputChan := generate(2, 3, 4)

	c1 := square(inputChan)
	c2 := square(inputChan)
	c3 := square(inputChan)

	for output := range merge(c1, c2, c3) {
		fmt.Println(output) // 작업 순서가 보장되지 않는다

	}

}
