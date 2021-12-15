package coding

import "fmt"

type status uint8

//other variations use white,gray,black as indicators of state
const (
	unprocessed status = iota
	processing
	processed
)

func dfs(adjList [][]int, statusArr []status, node int) error {
	statusArr[node] = processing
	for _, child := range adjList[node] {
		switch statusArr[child] {
		case processing:
			return fmt.Errorf("cycle detected")
		case unprocessed:
			err := dfs(adjList, statusArr, child)
			if err != nil { //dag is false
				return fmt.Errorf("cycle detected")
			}
		case processed:
			continue
		}
	}
	statusArr[node] = processed
	return nil
}

func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	res := true //assume dag is true initially
	// bi -> ai; bi is prereq to ai
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		prereq, nextCourse := prerequisites[i][1], prerequisites[i][0]
		graph[prereq] = append(graph[prereq], nextCourse)
	}
	// maintain a status slice as we process dfs calls
	statusArr := make([]status, numCourses)
	for i := range statusArr {
		if statusArr[i] != processed {
			if err := dfs(graph, statusArr, i); err != nil {
				return false //dag is false
			}
		}
	}
	return res

}

// BFS
func canFinishBFS(numCourses int, prerequisites [][]int) bool {

	graph := make(map[int][]int)

	dependencies := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		want := prerequisites[i][0]
		need := prerequisites[i][1]

		dependencies[want]++

		if _, ok := graph[need]; ok {
			graph[need] = append(graph[need], want)
		} else {
			graph[need] = append([]int{}, want)
		}
	}

	q := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		if dependencies[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		course := q[0]
		subClasses := graph[course]

		for i := 0; i < len(subClasses); i++ {
			dependencies[subClasses[i]]--

			if dependencies[subClasses[i]] == 0 {
				q = append(q, subClasses[i])
			}
		}
		q = q[1:]
	}

	for i := 0; i < numCourses; i++ {
		if dependencies[i] != 0 {
			return false
		}
	}

	return true
}
