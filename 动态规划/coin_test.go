package 动态规划

import "testing"

func TestCoin(t *testing.T) {
	t.Log(Coin([]int{1, 5, 6}, 11))
}

func TestCoinMemo(t *testing.T) {
	t.Log(CoinMemo([]int{1, 5, 6}, 11))
}
