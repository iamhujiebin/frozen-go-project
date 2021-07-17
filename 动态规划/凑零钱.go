package 动态规划

/*
	画棵树   amount = 10  coins = [2,3,4]
			C(10)
        2 /   3|    4\
       C(8)   C(7)    C(6)
      2/ 3| 4\ /|\   /  |  \
    C(6)C(5)C(4)
  ........................
C(0) C(-n)  ..  ..  .. .. ..
*/

// 凑零钱
// coins :零钱数组,每种面额假设有无数
// amount: 凑的面额总额
// min:需要用到的最小零钱个数
// 解法思路:
//  1. 凑到amount的最小硬币数 = 凑到amount-coin[i]的最小硬币数+1
// 时间复杂度: O(k^n),不一定是这个,但可以肯定是指数级别
// 有多少个节点,就是递归了多少次
// 同样可以用备忘录去解决重叠子问题
func Coin(coins []int, amount int) (minValue int) {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	minValue = INT_MAX
	for i := range coins {
		subCoin := Coin(coins, amount-coins[i])
		if subCoin < 0 {
			continue // 子问题无解
		}
		minValue = min(minValue, subCoin+1)
	}
	return minValue
}

func CoinMemo(coins []int, amount int) (minValue int) {
	memo := make(map[int]int)
	return coinMemo(coins, amount, memo)
}

func coinMemo(coins []int, amount int, memo map[int]int) (minValue int) {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if val, ok := memo[amount]; ok {
		return val
	}
	minValue = INT_MAX
	for i := range coins {
		subCoin := coinMemo(coins, amount-coins[i], memo)
		if subCoin < 0 {
			continue // 子问题无解
		}
		minValue = min(1+subCoin, minValue)
	}
	memo[amount] = minValue
	return memo[amount]
}
