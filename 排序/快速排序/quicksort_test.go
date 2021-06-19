package 快速排序

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{44, 2, 33, 4, 55, 6, 77, 5, 11}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}