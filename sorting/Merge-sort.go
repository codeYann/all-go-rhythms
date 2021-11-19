package sorting

func merge(a, b []int) []int {
	i, j := 0, 0
	max := make([]int, 0)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			max = append(max, a[i])
			i += 1
		} else {
			max = append(max, b[j])
			j += 1
		}
	}

	for i < len(a) {
		max = append(max, a[i])
		i += 1
	}

	for j < len(a) {
		max = append(max, b[j])
		j += 1
	}

	return max
}

func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	left := MergeSort(list[:len(list)/2])
	right := MergeSort(list[len(list)/2:])

	return merge(left, right)
}
