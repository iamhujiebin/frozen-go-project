package 堆排序

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{-1, 1, 4, 7, 2, 99, 5, 8, 33, 3, 6, 13,9, 0}
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
}
