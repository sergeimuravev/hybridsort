package hybridsort

import (
	_heap "container/heap"
	"sync"
	"sync/atomic"
)

// heap is a primitive thread-safe binary heap wrapper
type heap struct {
	mux    sync.Mutex
	buff   heapBuffer
	count  int32
	in     chan run
	out    chan run
	outReq chan bool
	done   chan bool
}

func makeHeap() *heap {
	h := heap{
		in:     make(chan run, 1),
		out:    make(chan run, 2),
		outReq: make(chan bool),
		done:   make(chan bool),
	}

	_heap.Init(&h.buff)

	go func() {
	loop:
		for {
			select {
			case <-h.done:
				break loop

			case <-h.outReq:
				if h.buff.Len() > 0 {
					h.out <- _heap.Pop(&h.buff).(run)
				}

			case r := <-h.in:
				_heap.Push(&h.buff, r)
			}
		}
	}()

	return &h
}

func (h *heap) Len() int {
	return int(h.count)
}

func (h *heap) Close() {
	h.done <- true
	close(h.in)
	close(h.out)
	close(h.outReq)
	close(h.done)
}

func (h *heap) Push(x run) {
	h.in <- x
	atomic.AddInt32(&h.count, 1)
}

func (h *heap) PushTwo(x, y run) {
	h.in <- x
	h.in <- y
	atomic.AddInt32(&h.count, 2)
}

func (h *heap) Pop() run {
	h.outReq <- true
	r := <-h.out
	atomic.AddInt32(&h.count, -1)
	return r
}

func (h *heap) PopTwo() (run, run) {
	h.outReq <- true
	x := <-h.out
	h.outReq <- true
	y := <-h.out
	atomic.AddInt32(&h.count, -2)
	return x, y
}
