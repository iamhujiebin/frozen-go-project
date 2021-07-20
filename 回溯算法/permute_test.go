package 回溯算法

import (
	"fmt"
	"testing"
)

func TestPermutation_Permute(t *testing.T) {
	per := new(Permutation)
	res := per.Permute([]int{1, 2, 3, 4, 5})
	fmt.Printf("res:%v\nlen:%v\n", res, len(res))
	res = per.Permute([]int{1, 2, 3})
	fmt.Printf("res:%v\nlen:%v\n", res, len(res))
}
