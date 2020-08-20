package hybridsort

import (
	_heap "container/heap"
	"sync"
)

// heap is a primitive thread-safe binary heap wrapper
type heap struct {
	mux  sync.Mutex
	buff heapBuffer
}

func makeHeap() *heap {
	h := heap{}
	_heap.Init(&h.buff)
	return &h
}

func (h *heap) Len() int {
	return h.buff.Len()
}

func (h *heap) Push(x run) {
	h.mux.Lock()
	defer h.mux.Unlock()
	_heap.Push(&h.buff, x)
}

func (h *heap) PushTwo(x, y run) {
	h.mux.Lock()
	defer h.mux.Unlock()
	_heap.Push(&h.buff, x)
	_heap.Push(&h.buff, y)
}

func (h *heap) Pop() run {
	h.mux.Lock()
	defer h.mux.Unlock()
	return _heap.Pop(&h.buff).(run)
}

func (h *heap) PopTwo() (run, run) {
	h.mux.Lock()
	defer h.mux.Unlock()
	return _heap.Pop(&h.buff).(run),
		_heap.Pop(&h.buff).(run)
}
