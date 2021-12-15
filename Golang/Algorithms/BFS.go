package algorithms

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

var q []*TreeNode

func bfsRecursive(node *TreeNode) {

	// 현재 노드 출력
	fmt.Println(node.Val)

	// 왼쪽 탐색
	if node.Left != nil {
		q = append(q, node.Left)

		// fmt.Printf("%v", q)
	}

	// 오른쪽 탐색
	if node.Right != nil {
		q = append(q, node.Right)
	}

	// 큐가 비면 종료
	if len(q) == 0 {
		return
	}

	// 다음 노드 순회
	next := q[0]
	q = q[1:]

	bfsRecursive(next)
}

func bfsLoop(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		l := len(q)
		var tmp []int
		for i := 0; i < l; i++ {
			node := q[0]
			fmt.Println(node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}

	return ans
}

func StartBFS() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	bfsRecursive(root)
	bfsLoop(root)
}
