package bubblesort

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	values := []int{5, 3, 2, 1, 4}
	values = BubbleSort(values)
	for index, value := range values {
		fmt.Println("index->", index + 1, "value->", value)
	}
	if values[0] != 1 {
		t.Error("failed")
	}
}