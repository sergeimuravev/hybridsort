package hybridsort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/sergeimuravev/hybridsort"
)

// Random
func BenchmarkStandard_1024_Random(b *testing.B) { benchmarkStandardSort(b, 1024, true) }
func BenchmarkHybrid_1024_Random(b *testing.B)   { benchmarkHybridSort(b, 1024, true) }
func BenchmarkStandard_32K_Random(b *testing.B)  { benchmarkStandardSort(b, 1024*32, true) }
func BenchmarkHybrid_32K_Random(b *testing.B)    { benchmarkHybridSort(b, 1024*32, true) }
func BenchmarkStandard_256K_Random(b *testing.B) { benchmarkStandardSort(b, 1024*256, true) }
func BenchmarkHybrid_256K_Random(b *testing.B)   { benchmarkHybridSort(b, 1024*256, true) }

// Sorted
func BenchmarkStandard_1024_Sorted(b *testing.B) { benchmarkStandardSort(b, 1024, false) }
func BenchmarkHybrid_1024_Sorted(b *testing.B)   { benchmarkHybridSort(b, 1024, false) }
func BenchmarkStandard_32K_Sorted(b *testing.B)  { benchmarkStandardSort(b, 1024*32, false) }
func BenchmarkHybrid_32K_Sorted(b *testing.B)    { benchmarkHybridSort(b, 1024*32, false) }
func BenchmarkStandard_256K_Sorted(b *testing.B) { benchmarkStandardSort(b, 1024*256, false) }
func BenchmarkHybrid_256K_Sorted(b *testing.B)   { benchmarkHybridSort(b, 1024*256, false) }

// Returns random test set
func createRandomTestSet(n int) hybridsort.Interface {
	data := IntSlice{}
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

// Returns partially sorted test set
func createSortedTestSet(n int) hybridsort.Interface {
	data := IntSlice{}
	for i := 0; i < n/128; i++ {
		for j := 0; j < 128; j++ {
			data = append(data, j)
		}
	}

	return data
}

// Benchmark helpers
func benchmarkHybridSort(b *testing.B, n int, random bool) {
	s := hybridsort.Sort{}.
		WithMinRunSize(32).
		WithMaxRunSize(255).
		WithDegreeOfParallelism(8)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var data hybridsort.Interface
		if random {
			data = createRandomTestSet(n)
		} else {
			data = createSortedTestSet(n)
		}
		b.StartTimer()

		s.Sort(data)

		b.StopTimer()
		if !sort.IsSorted(data) {
			panic("Not sorted.")
		}
		b.StartTimer()
	}
}

func benchmarkStandardSort(b *testing.B, n int, random bool) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var data hybridsort.Interface
		if random {
			data = createRandomTestSet(n)
		} else {
			data = createSortedTestSet(n)
		}
		b.StartTimer()

		sort.Stable(data)

		b.StopTimer()
		if !sort.IsSorted(data) {
			panic("Not sorted.")
		}
		b.StartTimer()
	}
}

// Test collection type
type IntSlice []int

func (s IntSlice) Len() int                            { return len(s) }
func (s IntSlice) Less(i, j int) bool                  { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)                       { s[i], s[j] = s[j], s[i] }
func (s IntSlice) ElementAt(i int) interface{}         { return s[i] }
func (s IntSlice) SetElementAt(i int, val interface{}) { s[i] = val.(int) }
