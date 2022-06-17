package main

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	NotMonitored int = iota
	MonitoredNoCam
	MonitoredWithCam
)

func minCameraCover(root *TreeNode) int {

	cams := 0
	top := dfs(root, &cams)

	if top == NotMonitored {
		return cams + 1
	} else {
		return cams
	}

}

func dfs(node *TreeNode, cams *int) int {
	if node == nil {
		return MonitoredNoCam
	}

	l, r := dfs(node.Left, cams), dfs(node.Right, cams)

	if l == MonitoredNoCam && r == MonitoredNoCam {
		return NotMonitored
	} else if l == NotMonitored || r == NotMonitored {
		*cams++
		return MonitoredWithCam
	} else {
		return MonitoredNoCam
	}
}
