package 计数排序

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := []int{1, 4, 7, 2, 5, 8, 3, -1, 6, -9, 99, 9}
	fmt.Println(arr)
	CountSort(arr)
	fmt.Println(arr)
}
