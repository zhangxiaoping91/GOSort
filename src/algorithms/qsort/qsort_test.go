package qsort

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	values := []int{5, 3, 2, 1, 4}
	values=QuickSort(values,0,len(values)-1)
	for index, value := range values {
		fmt.Println("index->", index + 1, "value->", value)
	}
}