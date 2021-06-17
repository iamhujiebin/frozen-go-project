package 选择排序

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	println("普通")
	arr := []int{1, 4, 3, 2, 1}
	fmt.Println(arr)
	SelectSort(arr)
	fmt.Println(arr)

	println("递归")
	arr = []int{1, 4, 3, 2, 1}
	fmt.Println(arr)
	SelectSort2(arr, 0)
	fmt.Println(arr)

	println("优化-递归")
	arr = []int{1, 3, 9, 2, 7, 1, 4, 3, 2, 1}
	fmt.Println(arr)
	SelectSort3(arr, 0, len(arr)-1)
	fmt.Println(arr)

	println("优化-for+while")
	arr = []int{1, 3, 9, 2, 7, 1, 4, 3, 2, 1}
	fmt.Println(arr)
	SelectSort4(arr)
	fmt.Println(arr)
}
