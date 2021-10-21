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

func (heap *MaxHeap) Insert(element int) {
	heap.elements = append(heap.elements, element)
	heap.maxHeapUp(len(heap.elements) - 1)
	heap.length += 1
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

func (heap *MaxHeap) maxHeapUp(index int) {
	parentPosition := getParentIndex(index)
	if heap.elements[index] > heap.elements[parentPosition] {
		heap.elements[parentPosition], heap.elements[index] = heap.elements[index], heap.elements[parentPosition]
		heap.maxHeapUp(parentPosition)
	}
}
