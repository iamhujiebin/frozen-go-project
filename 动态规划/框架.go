package 动态规划

/*
	动态规划的框架:
	特点:
	1.	重叠子问题:通过子问题,可以推算出原问题的答案(凑零钱、Fib、背包问题)
	2.	状态转移方程:如何通过子问题转为原问题 Fib: F(n) = F(n-1) + F(n-1)
	题型: 穷举+求最值

	解题思路:
	1.	明确base case F(0),F(1) = 0,1
	2.	明确dp函数: F(n) = F(n-1) + F(n-2)
	3.	明确选择: 背包问题中 dp[i][w] = max(dp[i-1][w],dp[i-1][w-wt[i-1G]]+val[i-1])
*/

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}