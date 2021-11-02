package classic

type circularLinked struct {
	head   *Node
	tail   *Node
	length int
}

func CreateCircularLinkedList() *circularLinked {
	return &circularLinked{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (circle circularLinked) GetHead() *Node {
	return circle.head
}

func (circle circularLinked) GetTail() *Node {
	return circle.tail
}

func (circle circularLinked) GetTailNext() *Node {
	return circle.tail.next
}

func (circle circularLinked) GetSize() int {
	return circle.length
}

func (circle *circularLinked) Append(element int) {
	node := &Node{Key: element, next: nil}
	if circle.head == nil {
		circle.head = node
		circle.tail = node
	} else {
		node.next = circle.head
		circle.tail.next = node
		circle.tail = node
	}
	circle.length += 1
}

func (circle *circularLinked) InsertFront(element int) {
	node := &Node{Key: element, next: nil}
	if circle.head == nil {
		circle.head = node
		circle.tail = node
	} else {
		node.next = circle.head
		circle.head = node
		circle.tail.next = circle.head
	}
	circle.length += 1
}
