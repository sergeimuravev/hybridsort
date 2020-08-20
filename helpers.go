package hybridsort

// insertionSort implements same-named sorting algorithm
func insertionSort(data Interface, left, right int) {
	var i, j int
	for i = left + 1; i <= right; i++ {
		for j = i; j > left && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
