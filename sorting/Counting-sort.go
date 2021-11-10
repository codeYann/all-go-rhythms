package sorting

func CountSort(list []int) []int {
	// Find max element
	max := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	// Counting occurrences
	newList := make([]int, max)
	for j := 0; j < len(list); j++ {
		newList[list[j]-1] += 1
	}

	// Sorting elements based on occurrences
	ordered := make([]int, 0)
	for i := 0; i < max; i++ {
		for j := 0; j < newList[i]; j++ {
			ordered = append(ordered, i+1)
		}
	}

	return ordered
}
