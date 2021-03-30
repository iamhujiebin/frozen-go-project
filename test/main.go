package main

import (
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"time"
)

var (
	num int32 = 0
)

func main() {
	var eg errgroup.Group
	eg.Go(func() error {
		atomic.AddInt32(&num, 1)
		return nil
	})
	eg.Go(func() error {
		atomic.AddInt32(&num, 1)
		return nil
	})
	eg.Go(func() error {
		atomic.AddInt32(&num, 1)
		return nil
	})
	eg.Wait()
	println(num)
	return
	println(time.Now().Unix(), "一秒前")
	timer := time.NewTimer(time.Second)
	<-timer.C
	println(time.Now().Unix(), "一秒后")
	timer.Stop()

	v := make(chan struct{})
	timer = time.AfterFunc(time.Second, func() {
		println(time.Now().Unix(), "再一秒后")
		v <- struct{}{}
	})
	<-v
	timer.Stop()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		done <- struct{}{}
	}()
Loop:
	for {
		select {
		case <-done:
			break Loop
		case <-ticker.C:
			println(time.Now().Unix(), "ticker")
		}
	}
}
