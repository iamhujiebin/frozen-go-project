package 基数排序

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	arr := []uint{111, 4, 171, 112, 511, 18, 1231, 16, 9, 1}
	fmt.Println(arr)
	RadixSort(arr)
	fmt.Println(arr)
}
