package queuestack

import (
	"fmt"
	"testing"
)

var cq CircleQueue
var circleQueue *QueueStack

func init() {
	circleQueue, _ = cq.Init(5, &Node{Data: 1}, &Node{Data: 2}, &Node{Data: 3})
}

func TestCircleQueue_Init(t *testing.T) {
	cq.Print(circleQueue)
}

func TestCircleQueue_Push(t *testing.T) {
	err := cq.Push(circleQueue, &Node{Data: 4})
	t.Log(err)
	err = cq.Push(circleQueue, &Node{Data: 5})
	t.Log(err)
	err = cq.Push(circleQueue, &Node{Data: 6})
	t.Log(err)
	cq.Print(circleQueue)
}

func TestCircleQueue_Pop(t *testing.T) {
	n := cq.Pop(circleQueue)
	t.Log(n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	t.Log(n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	t.Log(n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	t.Log(n)
	cq.Print(circleQueue)
}

func TestPushPop(t *testing.T) {
	cq.Print(circleQueue)
	cq.Push(circleQueue, &Node{Data: 4})
	n := cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	cq.Print(circleQueue)
	for !cq.IsEmpty(circleQueue) {
		n := cq.Pop(circleQueue)
		fmt.Printf("n:%v\n", n)
	}
	cq.Print(circleQueue)
	num := 100
	for !cq.IsFull(circleQueue) {
		cq.Push(circleQueue, &Node{Data: num})
		num++
	}
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	n = cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	cq.Print(circleQueue)
	n = cq.Pop(circleQueue)
	fmt.Printf("n:%v\n", n)
	cq.Print(circleQueue)
}
