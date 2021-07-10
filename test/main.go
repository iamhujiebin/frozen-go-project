package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), time.Millisecond*100)

func main() {
	testWait()
	//testAllow()
	//testReserve()
}

func testWait() {
	// 一秒r个，一次最多拿b个令牌
	limiter := rate.NewLimiter(10, 10)
	fmt.Println(limiter.Limit(), limiter.Burst())
	n := 0
	fmt.Printf("n:%v time:%v err:%v\n", n, time.Now(), nil)
	// 这里可以用于控制for循环的速度！
	for {
		// WaitN 就是消费n个token
		// 不消费token就不能往下走--阻塞
		// waitN(2) -> 一次消费两个token才能往下走
		// 加入用有timeout的ctx,会有err返回
		err := limiter.WaitN(context.Background(), 2)
		n++
		fmt.Printf("n:%v time:%v err:%v\n", n, time.Now(), err)
		if n == 100 {
			limiter.SetLimit(20)
		}
	}
}

func testAllow() {
	// 一秒10个
	limiter := rate.NewLimiter(10, 10)
	fmt.Println(limiter.Limit(), limiter.Burst())
	n := 0
	for {
		n++
		// 一次消耗2个令牌
		allow := limiter.AllowN(time.Now(), 2)
		if allow {
			fmt.Printf("n:%v time:%v err:%v\n", n, time.Now(), allow)
		} else {
			// 不阻塞
			// 这里的请求就丢掉了
		}
	}
}

func testReserve() {
	// 10秒1次
	limiter := rate.NewLimiter(0.1, 1)
	fmt.Println(limiter.Limit(), limiter.Burst())
	n := 0
	for {
		n++
		r := limiter.Reserve()
		// 除非设置burst失败,不然都是ok的
		if r.OK() {
			// 提前消耗(n)个token
			// 返回的reserve结构体告诉你Delay多久可以执行,0就是不用delay
			// 可以把token归还,那下一次获取的就不用叠加等待
			fmt.Printf("n:%v time:%v r:%+v\n", n, time.Now(), r.Delay())
		} else {
			fmt.Printf("should not come here")
		}
		time.Sleep(time.Second)
	}
}
