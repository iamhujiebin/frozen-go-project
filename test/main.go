package main

import (
	"context"
	"fmt"
	"frozen-go-project/utils/spawn"
	"frozen-go-project/utils/tls"
	"time"
)

func main() {
	tls.Set("jiebin", "mengyin")
	go func() {
		tls.Set("jiebin", "danxin")
		v, ok := tls.Get("jiebin")
		fmt.Printf("in:v:%+v,ok:%v\n", v, ok)
		tls.Flush()
	}()
	time.Sleep(time.Second)
	v, ok := tls.Get("jiebin")
	fmt.Printf("v:%+v,ok:%v\n", v, ok)

	go tls.For(context.WithValue(context.Background(), "mengyin", "dan"), func() {
		ctx, ok := tls.GetContext()
		fmt.Printf(" tls for ctx:%+v,ok:%v\n", ctx, ok)
	})()
	time.Sleep(time.Second)

	goSpawn := spawn.NewGoSpawn(10)
	for i := 0; i < 10; i++ {
		goSpawn.Go(tls.For(context.Background(), func() {
			k, v := tls.Get("jiebin")
			println(k, v)
		}))
	}
	time.Sleep(time.Second)
	println("end")
}
