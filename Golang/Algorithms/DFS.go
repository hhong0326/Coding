package algorithms

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func dfsRecursive(node *TreeNode) {

	if node == nil {
		return
	}

	fmt.Println(node.Val)

	if node.Left != nil {
		dfsRecursive(node.Left)
	}
	if node.Right != nil {
		dfsRecursive(node.Right)
	}
}

// using stack
func dfsLoop(node *TreeNode) {

	visit := []int{}
	stack := []*TreeNode{}
	stack = append(stack, node)

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		// fmt.Println(node.Val)
		stack = stack[:len(stack)-1]

		if len(visit) == 0 {
			visit = append(visit, node.Val)
			if node.Left != nil {

				stack = append(stack, node.Left)
			}
			if node.Right != nil {

				stack = append(stack, node.Right)
			}
		} else {
			for _, v := range visit {
				if node.Val != v {
					visit = append(visit, node.Val)
					if node.Left != nil {

						stack = append(stack, node.Left)
					}
					if node.Right != nil {

						stack = append(stack, node.Right)
					}
				}
			}

		}

	}

	fmt.Println(visit)
}

func StartDFS() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	dfsLoop(root)
}
