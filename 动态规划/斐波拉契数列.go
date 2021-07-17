package 动态规划

/*
	分析斐波拉契数列,用一个树去理解就很好理解。包括算法时间/空间复杂度分析和解题思路
			F(10)
         /        \
       F(9)       F(8)
      /   \      /    \
    F(8)  F(7)  F(7)  F(6)
  ........................
F(0) F(1)  ..  ..  .. .. ..
*/

// F(n) = F(n-1) + F(n-1)
// 时间复杂度 递归: O(2^n) * O(1) = O(2^n)
// 有多少个节点,就是递归了多少次
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// 备忘录模式,解决重叠子问题的重复结算
// 时间和空间复杂度都是O(n)
func FibMemo(n int) int {
	if n <= 1 {
		return n
	}
	memo := make(map[int]int)
	memo[0], memo[1] = 0, 1
	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

// 滚动数组
// 时间复杂度O(n),空间复杂度O(1)
func FibRollArr(n int) (r int) {
	if n <= 1 {
		return n
	}
	i, j := 0, 1
	for k := 2; k <= n; k++ {
		r = i + j
		i = j
		j = r
	}
	return r
}
