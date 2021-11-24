package sorting

func partition(A []int, start, end int) int {
	pivot := A[end]
	i := start
	for j := start; j < end; j++ {
		if A[j] <= pivot {
			A[j], A[i] = A[i], A[j]
			i += 1
		}
	}
	A[i], A[end] = A[end], A[i]
	return i + 1
}

func QuickSort(A []int, start, end int) {
	if start < end {
		p := partition(A, start, end)
		QuickSort(A, start, p-1)
		QuickSort(A, p+1, end)
	}
}
