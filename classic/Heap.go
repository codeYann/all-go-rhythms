package classic

type MaxHeap struct {
	elements []int
	length   int
	capacity int
}

func CreateMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		elements: make([]int, size),
		length:   0,
		capacity: size,
	}
}

func (heap *MaxHeap) GetElements() []int {
	return heap.elements
}

func (heap *MaxHeap) Size() int {
	return heap.length
}

func getParentIndex(index int) int {
	return (index) / 2
}

func leftChildren(index int) int {
	return (index * 2)
}

func rightChildren(index int) int {
	return (index * 2) + 1
}

func (heap *MaxHeap) Insert(element int) {
	if heap.length < heap.capacity {
		heap.length += 1
		heap.elements[heap.length] = element
		heap.maxHeapUp(heap.length)
	}
}

func (heap *MaxHeap) maxHeapUp(index int) {
	parentPosition := getParentIndex(index)
	if parentPosition > 0 {
		if heap.elements[index] > heap.elements[parentPosition] {
			heap.elements[parentPosition], heap.elements[index] = heap.elements[index], heap.elements[parentPosition]
			heap.maxHeapUp(parentPosition)
		}
	}
}

func (heap *MaxHeap) MaxHeapDown(i int) {
	l := heap.length
	left, right := leftChildren(i), rightChildren(i)
	choice := 0
	for left <= l {
		if left == l {
			choice = left
		} else if heap.elements[left] > heap.elements[right] {
			choice = left
		} else {
			choice = right
		}

		if heap.elements[i] < heap.elements[choice] {
			heap.elements[i], heap.elements[choice] = heap.elements[choice], heap.elements[i]
			i = choice
			left, right = leftChildren(i), rightChildren(i)
		} else {
			return
		}
	}
}

func (heap *MaxHeap) Pop() int {
	if heap.length > 0 {
		n := heap.length
		max := heap.elements[1]
		heap.elements[1] = heap.elements[n-1]
		heap.length = heap.length - 1
		heap.MaxHeapDown(1)
		return max
	}
	return -1
}
