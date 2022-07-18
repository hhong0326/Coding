package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	max := 1
	dfs(root, 1, &max)

	fmt.Println("max: ", max)
	return max
}

func dfs(node *TreeNode, depth int, max *int) {

	if node.Left == nil && node.Right == nil {
		if *max < depth {
			*max = depth
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, depth+1, max)
	}
	if node.Right != nil {
		dfs(node.Right, depth+1, max)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{}
	root.Left.Right = &TreeNode{}
	root.Right.Left = &TreeNode{}
	root.Right.Right = &TreeNode{}

	fmt.Println(root)
	maxDepth(root)
}
