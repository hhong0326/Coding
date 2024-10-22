type Data struct {
	Count int
	Symbol rune
}
type IntHeap []Data

func (h IntHeap)  Len() int           { return len(h) }
func (h IntHeap)  Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h IntHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(Data)) }
func (h *IntHeap) Pop() interface{}   {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	hp := IntHeap{}
	if a > 0 { hp = append(hp, Data{Count: a, Symbol: 'a'}) }
	if b > 0 { hp = append(hp, Data{Count: b, Symbol: 'b'}) }
	if c > 0 { hp = append(hp, Data{Count: c, Symbol: 'c'}) }

	heap.Init(&hp)
	res, last, prev := "", 'd', 'd'
	for len(hp) > 0 {
			cur := heap.Pop(&hp).(Data)
			if cur.Symbol == last && prev == last {
					if len(hp) > 0 {
							newCur := heap.Pop(&hp).(Data)
							res += string(newCur.Symbol)
							newCur.Count--
							if newCur.Count > 0 {
									heap.Push(&hp, Data{ Count: newCur.Count, Symbol: newCur.Symbol })
							}
							heap.Push(&hp, Data{ Count: cur.Count, Symbol: cur.Symbol })
							prev  = last
							last = newCur.Symbol
					}
			} else {
					res += string(cur.Symbol)
					cur.Count--
					if cur.Count > 0 {
							heap.Push(&hp,  Data{ Count: cur.Count, Symbol: cur.Symbol })
					}
					prev = last
					last = cur.Symbol
			}
	}
	return res
}