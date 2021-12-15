package coding

// BFS

func canReach(arr []int, start int) bool {

	q := []int{start}
	visited := make([]bool, len(arr))
	visited[start] = true

	for len(q) > 0 {

		s := q[0]
		q = q[1:]

		if arr[s] == 0 {
			return true
		}

		if s-arr[s] >= 0 && !visited[s-arr[s]] {
			q = append(q, s-arr[s])
			visited[s-arr[s]] = true
		}

		if s+arr[s] < len(arr) && !visited[s+arr[s]] {
			q = append(q, s+arr[s])
			visited[s+arr[s]] = true

		}

	}

	return false
}
