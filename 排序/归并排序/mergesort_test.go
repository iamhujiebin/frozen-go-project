package 归并排序

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	a1 := []int{1, 3, 5}
	a2 := []int{2, 4, 6}
	m := Merge(a1, a2)
	fmt.Println(m)
}

func TestMerge2(t *testing.T) {
	a1 := []int{1, 4, 7}
	a2 := []int{2, 5, 8}
	a3 := []int{3, 6, 9}
	m := Merge(Merge(a1, a2), a3)
	fmt.Println(m)
}

func TestMergeSort(t *testing.T) {
	arr := []int{-1, 2, 9, 99, 44, 1, 3, 2, 84, 16}
	fmt.Println(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestMergeSort2(t *testing.T) {
	//arr := []int{-1, 2, 9, 99, 44, 1, 3, 2, 84, 16}
	arr := []int{1, 4, 7, 2, 5, 8}
	fmt.Println(arr)
	MergeSort2(arr)
	fmt.Println(arr)
}

// 注意,如果要改变入参的[]int,需要用指针
func test(m *[]int) {
	*m = append(*m, 1)
	*m = append(*m, 1)
	*m = append(*m, 1)
}
