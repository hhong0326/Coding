package coding

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	arr := bfsZigzag(root)
	for i := 0; i < len(arr); i++ {
		if i%2 != 0 {
			reverseArr(arr[i])
		}
	}
	return arr
}

func reverseArr(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1] = arr[len(arr)-1-i], arr[i]
	}
}

func bfsZigzag(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var results [][]int
	var q []*TreeNode

	q = append(q, root)

	for len(q) > 0 {
		l := len(q)
		var tmp []int

		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			tmp = append(tmp, node.Val)
		}

		results = append(results, tmp)
	}

	return results
}
