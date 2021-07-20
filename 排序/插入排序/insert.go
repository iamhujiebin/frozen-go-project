package 插入排序

import (
	"fmt"
	"strconv"
)

func InsertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		index := i
		value := arr[i] // 找扑克牌位置
		for j := i; j >= 0; j-- {
			if arr[j] > arr[i] {
				index = j
			}
		}
		if index != i {
			copy(arr[index+1:i+1], arr[index:i+1]) // golang语法,左闭右开
			arr[index] = value
		}
	}
}

func test() {
	n := 13
	var res []string
	for i := 1; i <= n; i++ {
		res = insert(res, strconv.Itoa(i))
	}
	fmt.Println(res)
}

func insert(arr []string, data string) []string {
	arr = append(arr, data) // 这一步很关键,往数组中加一个元素,然后排序
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] < arr[i] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}
	return arr
}
