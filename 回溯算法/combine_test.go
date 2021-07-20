package 回溯算法

import (
	"fmt"
	"testing"
)

func TestCombination_Combine(t *testing.T) {
	com := new(Combination)
	res := com.Combine([]int{1, 2, 3}, 2)
	fmt.Printf("res:%v\n len:%v", res, len(res))
	res = com.Combine([]int{2, 9, 11, 4, 5}, 2)
	fmt.Printf("res:%v\n len:%v", res, len(res))
}
