package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 13
	var res []string
	for i := 1; i <= n; i++ {
		res = insert(res, strconv.Itoa(i))
	}
	fmt.Println(res)
}

func insert(arr []string, data string) []string {
	arr = append(arr, data)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] < arr[i] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}
	return arr
}
