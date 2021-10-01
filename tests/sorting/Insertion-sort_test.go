package tests

import (
	"testing"

	s "github.com/codeYann/all-go-rhythms/sorting"
)

func TestInsertionSort(t *testing.T) {
	slice := []int{10, 16, 13, 18, 32, 56, 1, 2, 0, 8, 9, 17, 11, 3, 28, 26, 44, 64}
	min, max := slice[8], slice[17]
	s.InsertionSort(slice)

	if slice[0] != min {
		t.Error("Expected:", min, "Got:", slice[0])
	} else {
		t.Log("Success")
	}

	if slice[17] != max {
		t.Error("Expected:", max, "Got:", slice[17])
	} else {
		t.Log("Success")
	}
}

func TestRepeatedElements(t *testing.T) {
	slice := []int{10, 10, 1, 1, 1, 2, 2, 3, 6, 27, 28, 64, 128, 13}
	min, max := slice[2], slice[12]
	s.InsertionSort(slice)

	if slice[0] != min {
		t.Error("Expected:", min, "Got:", slice[0])
	}

	if slice[13] != max {
		t.Error("Expected:", max, "Got:", slice[13])
	}
}
