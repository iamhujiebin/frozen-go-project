package 观察者Observer

import (
	"fmt"
	"sync"
)

// 观察者模式
// 生产者/消费者中的消费者,有notifyCallback即可
// 三个对象,观察者:Observer 主题:Subject 生产者:Publish
type Event struct {
	state string
}

type Observer interface {
	NotifyCallback(event Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type ConcreteObserver struct {
	Name string
}

func (c *ConcreteObserver) NotifyCallback(event Event) {
	fmt.Printf("notify call back:%s\n", event.state)
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{observers: sync.Map{}}
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{Name: name}
}

type ConcreteSubject struct {
	observers sync.Map
}

func (c *ConcreteSubject) AddListener(observer Observer) {
	c.observers.Store(observer, struct{}{})
}

func (c *ConcreteSubject) RemoveListener(observer Observer) {
	c.observers.Delete(observer)
}

func (c *ConcreteSubject) Notify(event Event) {
	c.observers.Range(func(key, value interface{}) bool {
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}
