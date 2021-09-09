package classic

import (
	"testing"
)

func TestCreateStack(t *testing.T) {
	stack := CreateStack()

	if stack.IsEmpty() != 0 {
		t.Error("Expected:", 0, "Got:", stack.Size())
	} else {
		t.Log("Success on create a stack")
	}

	if stack.Top() != nil {
		t.Error("Expected:", nil, "Got:", stack.Top())
	} else {
		t.Log("Success on create a stack")
	}

}

func TestPush(t *testing.T) {
	stack := CreateStack()
	values := []int{99, 75, 971, 347, 313, 317, 81, 40}

	for _, v := range values {
		stack.Push(v)
	}

	if stack.Top().Key != values[7] {
		t.Error("Expected:", values[7], "Got:", stack.Top().Key)
		println(stack.Top().Key, values[7])
	} else {
		t.Log("Success on push values")
	}

	if stack.Size() != 8 {
		t.Error("Expected:", 8, "Got:", stack.Size())
	}
}

func TestPop(t *testing.T) {
	stack := CreateStack()
	values := []int{99, 75, 971, 347, 313, 317, 81, 40}

	for _, v := range values {
		stack.Push(v)
	}

	for j := 0; j < 4; j++ {
		stack.Pop()
	}

	if stack.Top().Key != values[3] {
		t.Error("Expected:", values[3], "Got:", stack.Top().Key)
	} else {
		t.Log("Success on pop values")
	}

	if stack.Size() != 4 {
		t.Error("Expected:", 4, "Got:", stack.Size())
	} else {
		t.Log("Success on pop values")
	}
}
