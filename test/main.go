package main

import (
	"context"
	"fmt"
	"frozen-go-project/cmd/tls"
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
}
