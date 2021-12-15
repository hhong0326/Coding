package algorithms

import "fmt"

func binarySearchTree(nums []int, target int) int { //index

	//left, right := pivot, pivot-1+n
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid%n]

		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid % n
		}
	}

	return -1
}

func StartBST() {
	l := []int{1, 25, 30, 33, 40, 45, 57, 59, 70, 99}

	index := binarySearchTree(l, 40)
	fmt.Println(l[index])
}
