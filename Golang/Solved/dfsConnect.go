package coding

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}

	var prev *Node
	for len(q) > 0 {

		tmp := []*Node{}
		for i := 0; i < len(q); i++ {
			node := q[i]

			if prev != nil {
				prev.Next = node
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}

			prev = node
		}

		q, prev = tmp, nil // last node.Next -> nil
	}

	return root
}
