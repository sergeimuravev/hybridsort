package hybridsort

import (
	"runtime"
	"sync"
)

// Sort struct contains settings used by sorting procedure.
type Sort struct {
	minRunSize int
	maxRunSize int
	dop        int
}

// Sort implements hybrid stable sorting algorithm.
func (sort *Sort) Sort(data Interface) {
	// Ensure correct settings applied
	sort.checkSettings()

	// Fetch runs
	ch := make(chan run, sort.dop)
	var n, length int = 0, data.Len()
	go func() {
		defer close(ch)
		for i := 0; i < length; {
			left, right := sort.nextRun(data, i)
			r := run{Start: left, End: right, Seq: n}
			i = right + 1
			n++
			ch <- r
		}
	}()

	// Sort runs
	heap := makeHeap()
	wg := &sync.WaitGroup{}
	for run := range ch {
		wg.Add(1)
		r := run
		go func() {
			defer wg.Done()
			insertionSort(data, r.Start, r.End)
			heap.Push(r)
		}()
	}
	wg.Wait()

	// Select runs to merge
	semaphore := make(chan bool, sort.dop)
	len := 0
	for {
		len = heap.Len()
		if len > 1 {
			var x, y run = heap.PopTwo()
			if y.Start-x.End != 1 {
				// Stop and replay non-contiguous runs
				wg.Wait()
				x.Seq++
				heap.PushTwo(x, y)
				continue
			}

			semaphore <- false
			wg.Add(1)
			go func() {
				// x | y
				defer func() {
					wg.Done()
					<-semaphore
				}()
				symMerge(data, x.Start, y.Start, y.End+1)
				x.End = y.End
				x.Seq += n
				heap.Push(x)
			}()
		} else if len == 1 {
			// Test stop
			wg.Wait()
			if heap.Len() == 1 {
				close(semaphore)
				heap.Close()
				break
			}
		}
	}
}

// checkSettings provides defaults if settings have not been assigned.
func (sort *Sort) checkSettings() {
	if sort.minRunSize == 0 {
		sort.minRunSize = 32
	}

	if sort.maxRunSize == 0 {
		sort.maxRunSize = 255
	}

	if sort.dop == 0 {
		sort.dop = runtime.NumCPU()
	}
}

// nextRun selects the next run (chunk) of elements to process. It returns the first the last index of the run.
// The run size is limited by [min, max]. Selection is trying to find pre-sorted sequence of elements.
// If no pre-sorted sequence is found, selection is filled until [min] number of elements is provided.
func (sort *Sort) nextRun(data Interface, offset int) (int, int) {
	var i, last int = offset, data.Len() - 1
	for i < last &&
		(i-offset+1) <= sort.maxRunSize &&
		data.Less(i, i+1) {
		i++
	}

	if (i-offset+1) < sort.minRunSize &&
		i < last {
		i += sort.minRunSize - (i - offset + 1)

		if i > last {
			i = last
		}
	}

	return offset, i
}
