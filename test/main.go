package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 99, 0, 99, 99, 99, 99}
	for i, v := range arr {
		if v == 2 || v == 5 {
			arr = append(arr[0:i], arr[i+1:]...)
		}
	}
	fmt.Println(arr)
}
