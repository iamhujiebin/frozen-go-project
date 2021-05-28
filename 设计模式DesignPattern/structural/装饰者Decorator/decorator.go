package 装饰者Decorator

import (
	"fmt"
	"math"
	"time"
)

// 装饰者模式
// Wrap,在实际的func中,多加一个侧面,如统计时间
func WrapCalCost(fun piFunc) piFunc {
	return func(i int) (result float64) {
		defer func(t time.Time) {
			fmt.Printf("cal func cost:%v,result:%v\n", time.Since(t), result)
		}(time.Now())
		return fun(i)
	}
}

type piFunc func(int) float64

// 实际的方法
func Pi(n int) (result float64) {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(i))
	}
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return
}
