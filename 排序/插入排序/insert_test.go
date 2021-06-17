package 插入排序

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	println("普通")
	arr := []int{1, 2, 5, 9, 7, 2, 4, 3, 2, 1}
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)

	arr = []int{3, 4, 2, 1, 7, 6, 9, 11, 12, 5}
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)
}
