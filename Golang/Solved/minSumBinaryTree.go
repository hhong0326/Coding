package main

import (
	"fmt"
	"sort"
	"time"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func dfsLoop(root *Node) int {

	s := []*Node{root}

	var arr []int
	var sum int
	for len(s) > 0 {

		node := s[len(s)-1]
		s = s[:len(s)-1]

		sum += node.Value
		if node.Left == nil && node.Right == nil {
			arr = append(arr, sum)
			sum = root.Value
			continue
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
		// fmt.Println(node.Value)
	}

	sort.Ints(arr)
	return arr[0]
}

func main() {

	root := &Node{Value: 10}
	root.Left = &Node{Value: 5}
	root.Right = &Node{Value: 5}
	root.Left.Right = &Node{Value: 2}
	root.Right.Right = &Node{Value: 1}
	root.Right.Right.Left = &Node{Value: -1}

	s3 := time.Now()
	fmt.Println(dfsLoop(root), time.Since(s3))
}
