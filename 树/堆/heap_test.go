package 堆

import (
	"fmt"
	"testing"
)

type ConcreteElement struct {
	priority int
	Data     interface{}
}

func (c ConcreteElement) Priority() int {
	return c.priority
}

func (c ConcreteElement) Bigger(element Element) bool {
	return c.Priority() > element.Priority()
}

func TestHeap_InitHeap(t *testing.T) {
	h := new(Heap)
	h.InitHeap(10)
	h.Push(ConcreteElement{priority: 1, Data: 1111})
	h.Push(ConcreteElement{priority: 2, Data: 21111})
	h.Push(ConcreteElement{priority: 3, Data: 3111})
	h.Push(ConcreteElement{priority: 4, Data: 4111})
	h.Push(ConcreteElement{priority: 5, Data: 51111})
	h.Push(ConcreteElement{priority: 6, Data: 6111})
	h.Push(ConcreteElement{priority: 7, Data: 711})
	h.Push(ConcreteElement{priority: 8, Data: 81})
	h.Push(ConcreteElement{priority: 9, Data: 91})
	h.Push(ConcreteElement{priority: 10, Data: 10})
	h.Push(ConcreteElement{priority: 19, Data: 191111})
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	println(h)
}
