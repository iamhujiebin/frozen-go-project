package 装饰者Decorator

import "testing"

func TestWrapCalCost(t *testing.T) {
	wrapPi := WrapCalCost(Pi)
	wrapPi(10)
	wrapPi(1000)
	wrapPi(2000)
	wrapPi(20000)
}
