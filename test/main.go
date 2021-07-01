package main

import "fmt"

type query func(string) string

func main() {
	ret := exec("111", func(s string) string {
		return s + " func1"
	}, func(s string) string {
		return s + " func2"
	}, func(s string) string {
		return s + " func3"
	})
	fmt.Println(ret)
}

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i := range vs {
		go fn(i)
	}
	return <-ch
}
