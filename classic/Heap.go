package classic

type MaxHeap struct {
	elements []int
	length   int
}

func CreateMaxHeap() *MaxHeap {
	return &MaxHeap{
		elements: make([]int, 0),
		length:   0,
	}
}

func (heap *MaxHeap) GetElements() []int {
	return heap.elements
}

func (heap *MaxHeap) Size() int {
	return heap.length
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func leftChildren(index int) int {
	return (index * 2) + 1
}

func rightChildren(index int) int {
	return (index * 2) + 2
}

func (heap *MaxHeap) Insert(element int) {
	heap.elements = append(heap.elements, element)
	heap.maxHeapUp(len(heap.elements) - 1)
	heap.length += 1
}

func (heap *MaxHeap) maxHeapUp(index int) {
	parentPosition := getParentIndex(index)
	if heap.elements[index] > heap.elements[parentPosition] {
		heap.elements[parentPosition], heap.elements[index] = heap.elements[index], heap.elements[parentPosition]
		heap.maxHeapUp(parentPosition)
	}
}

func (heap *MaxHeap) MaxHeapDown(i int) {
	l := len(heap.elements) - 1
	j := 0
	left, right := leftChildren(i), rightChildren(i)
	for left <= l {
		if left == l {
			j = left
		} else if heap.elements[left] > heap.elements[i] {
			j = left
		} else {
			j = right
		}

		if heap.elements[i] < heap.elements[j] {
			heap.elements[i], heap.elements[j] = heap.elements[j], heap.elements[i]
			i = j
			left, right = leftChildren(i), rightChildren(i)
		} else {
			return
		}
	}
}

func (heap *MaxHeap) Pop() int {
	MaxValue := 0
	if heap.length >= 1 {
		MaxValue = heap.elements[0]
		heap.length -= 1
		heap.elements[0] = heap.elements[len(heap.elements)-1]
		heap.elements = heap.elements[:len(heap.elements)-1]
		heap.MaxHeapDown(0)
	}
	return MaxValue
}
