package 动态规划

import "testing"

func TestKnapsack(t *testing.T) {
	val := []int{4, 5, 13}
	wt := []int{1, 3, 4}
	W := 4
	t.Log(Knapsack(wt, val, W))
}
