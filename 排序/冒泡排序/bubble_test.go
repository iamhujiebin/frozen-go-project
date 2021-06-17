package 冒泡排序

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 4, 3, 2, 1}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)

	arr = []int{1, 4, 3, 2, 1}
	fmt.Println(arr)
	BubbleSort2(arr, len(arr))
	fmt.Println(arr)
}
