package hybridsort

// heapBuffer is an internal binary heap in-memory storage
type heapBuffer []run

func (h heapBuffer) Len() int           { return len(h) }
func (h heapBuffer) Less(i, j int) bool { return h[i].Seq < h[j].Seq }
func (h heapBuffer) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapBuffer) Push(x interface{}) {
	*h = append(*h, x.(run))
}

func (h *heapBuffer) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
