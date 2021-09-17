package classic

import "errors"

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *LinkedList) PrintList() {
	temp := l.head
	for temp != nil {
		println(temp.Key)
		temp = temp.next
	}
}

func (l *LinkedList) Size() int {
	return l.length
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) AddBack(value int) {
	node := &Node{Key: value, next: nil}

	if l.head == nil || l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length += 1
}

func (l *LinkedList) AddFront(value int) {
	node := &Node{Key: value, next: nil}
	if l.head == nil || l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.length += 1
}

func (l *LinkedList) RemoveBack() (*Node, error) {
	if l.head == nil || l.tail == nil {
		return nil, errors.New("List is empty!")
	} else {
		temp := l.head
		removedNode := l.tail

		for temp != nil {
			if temp.next == l.tail {
				break
			}
			temp = temp.next
		}

		temp.next = nil
		l.tail = temp
		l.length -= 1
		return removedNode, nil
	}
}

func (l *LinkedList) RemoveFront() (*Node, error) {
	if l.head == nil || l.tail == nil {
		return nil, errors.New("List is empty")
	} else {
		temp := l.head
		l.head = l.head.next
		l.length -= 1
		return temp, nil
	}
}
