package main

import "fmt"

func main() {

	for i := range fib(1000) {
		fmt.Println("피보나치 수: ", i)
	}
}

func fib(n int) chan int {
	out := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
		close(out)
	}()
	return out
}
