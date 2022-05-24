package main

import (
	"container/heap"
	"fmt"
)

// Heap은 그래프의 트리 구조 중 하나로 '우선순의 큐'를 구현할 때 사용된다.
// Heap을 표현하는 트리 구조에서는 각 정점을 '노드'라고 부른다.
// 자식 노드의 숫자는 반드시 부모 숫자보다 커야한다는 규칙이 있다.
// 따라서 가장 높은 노드는 가장 작은 숫자가 들어 있고
// 규칙을 지키기 위해 가장 아래에 있는 왼쪽 노드부터 값을 채운다.

// Push
// 가장 아래 부터 채운다.
// 만약 삽입 데이터가 부모보다 작다면 부모와 자식을 교환한다.
// 교환이 일어나지 않을 때 까지 반복한다.

// Pop
// 가장 위의 숫자가 추출된다.
// 가장 위가 없어졌으므로 트리를 다시 정리한다.
// 가장 후미에 있는 숫자를 가장 높은 루트로 이동시키고
// 부모 숫자보다 자식 숫자가 작은 경우
// 자식의 좌우 숫자중 더 작은 쪽과 교환하는 방식으로
// 트리를 재정리한다.

// 시간 복잡도
// Heap은 최솟값이 위에 있으므로 추출시 O(1)의 시간으로 데이터를 추출할 수 있다.
// 추출 후 Heap을 정리할 때는 가장 후미의 데이터를 위로 가져와 자식과 비교하며 아래 방향으로 진행된다.
// 즉, 계산 시간은 트리의 높이와 비례하고 데이터 수가 n이면
// Heap 형성 조건에 따라 트리의 높이는 log2n이므로 재구축 시간은 O(logn)이 된다.

// 활용
// 관리하는 데이터 중 최솟값을 자주 추출해야하는 경우 Heap을 사용

type IntHeap []int

// sort.Interface
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// heap.Interface
func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil
	*h = old[:n-1]
	return item
}
func main() {

	// golang의 container/heap이라는 내장 라이브러리를 활용한 minHeap 구현
	h := &IntHeap{2, 1, 5}

	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

// Reference
// https://aidanbae.github.io/code/golang/heap/
