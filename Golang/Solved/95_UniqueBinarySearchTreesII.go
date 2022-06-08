package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return tracker(1, n)
}

func tracker(s, e int) []*TreeNode {
	if s > e {
		return []*TreeNode{nil}
	}
	if s == e {
		return []*TreeNode{{
			Val:   s,
			Left:  nil,
			Right: nil,
		}}
	}

	res := []*TreeNode{}
	for i := s; i <= e; i++ {

		left := tracker(s, i-1)
		right := tracker(i+1, e)

		for _, l := range left {
			for _, r := range right {
				head := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				res = append(res, head)
			}
		}

	}
	return res
}
