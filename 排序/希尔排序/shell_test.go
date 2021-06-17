package 希尔排序

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{1, 9, 5, 3, 6, 7, 2, 11, 3, 9, 7, 4}
	fmt.Println(arr)
	ShellSort(arr)
	fmt.Println(arr)
}
