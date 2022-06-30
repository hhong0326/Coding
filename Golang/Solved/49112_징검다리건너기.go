package main

import "fmt"

var min = 10000

func solution(arr []int) {

	if len(arr) <= 3 {
		fmt.Println(0)
		return
	}

	back(arr, -1, 0)
	fmt.Println("min: ", min)
}

func back(arr []int, idx int, sum int) {

	if idx >= len(arr)-1 {
		if min > sum {
			min = sum
		}
		return
	}

	if idx < len(arr)-3 {
		back(arr, idx+1, sum+arr[idx+1])
		back(arr, idx+2, sum+arr[idx+2])
		back(arr, idx+3, sum+arr[idx+3])
	} else {
		if min > sum {
			min = sum
		}
		return
	}

}

func main() {
	solution([]int{7, 8, 9, 10, 9, 8, 7, 6, 5, 4})
}
