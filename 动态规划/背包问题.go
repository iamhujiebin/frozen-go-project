package 动态规划

// 背包问题
// wt:物品的数量数组  例如:[1,2,3]
// val:物品的价值数组 例如:[3,4,6]
// W:背包的容量
// 求:背包能装下的最大价值
// 解题思路:
// 	1. dp数组,dp[i][w]:背包容量为w放下前i个物品的最大价值
//  2. dp[0][w] = dp[i][0] = 0
//  3. 状态转移方程(子问题) : dp[i][w] = max(dp[i-1][w],dp[i-1][w-wt[i]+val[i])
//  3.1 背包的最大容量=放下i的最大容量或者不放下i的最大容量
func Knapsack(wt, val []int, W int) (maxValue int) {
	N := len(val)
	// 初始化dp数组 dp[N+1][W+1]
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	// 穷举wt/val
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
				continue
			}
			dp[i][w] = max(dp[i-1][w], dp[i-1][w-wt[i-1]]+val[i-1])
		}
	}
	return dp[N][W]
}