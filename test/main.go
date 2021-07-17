package main

import "fmt"

const W = 3

func main() {
	val := []int{4, 5, 8}
	wt := []int{1, 2, 3}
	N := len(val)
	dp := make([][W + 1]int, N+1)
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 越界,那只能不放了
				dp[i][w] = dp[i-1][w]
				continue
			}
			// 要么不放i多,要么放i多
			dp[i][w] = max(dp[i-1][w], val[i-1]+dp[i-1][w-wt[i-1]])
		}
	}
	fmt.Println(dp[N][W])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
