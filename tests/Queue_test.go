package tests

import (
	"testing"
	Queue "github.com/codeYann/all-go-rhythms/classic"
)

func checkNodes(node *Queue.Node, t *testing.T) {
	if node != nil {
		t.Error("Expected:", nil, "Got:", node)
	} else {
		t.Log("Success")
	}
}

func TestCreateQueue(t *testing.T) {
	queue := Queue.CreateQueue()

	if queue.IsEmpty() != 0 {
		t.Error("Expected:", 0, "Got:", queue.Size())
	} else {
		t.Log("Success on create a queue")
	}

	frontal, _ := queue.Frontal()
	back, _ := queue.Back()

	checkNodes(frontal, t)
	checkNodes(back, t)
}

func TestEnqueue(t *testing.T) {
	queue := Queue.CreateQueue()
	setNumbers := []int{10, 19, 72, 84, 110, 192}

	for i := 0; i < len(setNumbers); i++ {
		queue.Enqueue(setNumbers[i])
	}

	front, _ := queue.Frontal()
	back, _ := queue.Back()

	if front.Key != 10 {
		t.Error("Expected:", setNumbers[0], "got:", front.Key)
	} else {
		t.Log("Success on enqueue!")
	}

	if back.Key != 192 {
		t.Error("Expected:", setNumbers[5], "got:", back.Key)
	} else {
		t.Log("Success on enqueue!")
	}
}

func TestDequeue(t *testing.T) {
	queue := Queue.CreateQueue()

	setNumbers := []int{33, 74, 195, 157, 133, 1024}
	for i := 0; i < len(setNumbers); i++ {
		queue.Enqueue(setNumbers[i])
	}

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	front, _ := queue.Frontal()
	back, _ := queue.Back()

	if front.Key != setNumbers[3] {
		t.Error("Expected:", setNumbers[3], "got:", front.Key)
	} else {
		t.Log("Success on dequeue!")
	}

	if back.Key != setNumbers[5] {
		t.Error("Expected:", setNumbers[5], "got:", back.Key)
	} else {
		t.Log("Success on dequeue!")
	}
}
