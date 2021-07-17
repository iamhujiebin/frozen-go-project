package 动态规划

import "testing"

func TestFib(t *testing.T) {
	t.Log(Fib(10))
}

func TestFibMemo(t *testing.T) {
	t.Log(FibMemo(10))
}

func TestFibRollArr(t *testing.T) {
	t.Log(FibRollArr(10))
}
