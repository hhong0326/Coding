package main

import "container/heap"

type Item struct {
	request, process, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	prev := pq[i]
	next := pq[j]

	if prev.process == next.process {
		return prev.request < next.request
	}

	return prev.process < next.process
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func solution(jobs [][]int) int {
	jobsLen := len(jobs)
	jobsQueue := make(PriorityQueue, jobsLen)

	jobsList := make([]Item, jobsLen)

	time := 0
	sum := 0

	for i := 0; i < jobsLen; i++ {
		job := jobs[i]

		jobsQueue[i] = &Item{
			request: job[0],
			process: job[1],
			index:   i,
		}
	}

	heap.Init(&jobsQueue)

	for i := 0; i < jobsLen; i++ {
		job := heap.Pop(&jobsQueue).(*Item)

		jobsList = append(jobsList, *job)
	}

	for len(jobsList) > 0 {
		for i := 0; i < len(jobsList); i++ {
			job := jobsList[i]
			if time >= job.request {
				time += job.process
				sum += time - job.request
				jobsList = jobsList[:i+copy(jobsList[i:], jobsList[i+1:])]
				break
			}

			if i == len(jobsList)-1 {
				time += 1
			}
		}
	}

	return sum / jobsLen
}

// func solution(jobs [][]int) int {

// 	sort.Slice(jobs, func(i, j int) bool {
// 		return jobs[i][0] < jobs[j][0]
// 	})

// 	// start + spend time
// 	// if 작업
// 	// work finish time(t + spend time) + spend time

// 	answer := 0
// 	idx, e := 0, 0
// 	count := 0

// 	// queue
// 	q := [][]int{}

// 	// priority queue

// 	// 모든 job이 수행될 때까지 반복
// 	for count < len(jobs) {

// 		// 작업 중에 수행시간이 다가와 대기가 필요한 것을 큐에 넣는다
// 		for idx < len(jobs) && jobs[idx][0] <= e {
// 			q = append(q, jobs[idx])
// 			idx++
// 		}

// 		// 작업 중에는 수행할 게 없다면
// 		// 시간 업데이트
// 		if len(q) == 0 {
// 			e = jobs[idx][0]
// 		} else {
// 			// 큐 안에서 가장 수행시간이 짧은 것부터 수행
// 			sort.Slice(q, func(i, j int) bool {
// 				return q[i][1] < q[j][1]
// 			})

// 			job := q[0]
// 			q = q[1:]

// 			answer += job[1] + e - job[0]
// 			e += job[1]
// 			count++
// 		}
// 	}
// 	return int(math.Floor(float64(answer) / float64(len(jobs))))
// }

func main() {
	solution([][]int{{0, 3}, {1, 9}, {2, 6}})
}

// Reference
// https://codevang.tistory.com/316
