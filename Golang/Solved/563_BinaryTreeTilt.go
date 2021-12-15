package coding

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS, depending on the relative order of visit among the node and its children nodes.

func findTilt(root *TreeNode) int {
	var tilt int

	sumTilt(root, &tilt)
	return tilt
}

func sumTilt(node *TreeNode, tilt *int) int {
	if node == nil {
		return 0
	}

	lSum, rSum := sumTilt(node.Left, tilt), sumTilt(node.Right, tilt)
	*tilt += int(math.Abs(float64(lSum - rSum)))

	return node.Val + lSum + rSum
}
