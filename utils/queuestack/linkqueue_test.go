package queuestack

import "testing"

var linkqueue *QueueStack
var lq LinkQueue

func init() {
	linkqueue, _ = lq.Init(0, 1, 2, 3)
}

func TestLinkQueue_Init(t *testing.T) {
	lq.Print(linkqueue)
}

func TestLinkQueue_Push(t *testing.T) {
	_ = lq.Push(linkqueue, 4)
	lq.Print(linkqueue)
}

func TestLinkQueue_Pop(t *testing.T) {
	node := lq.Pop(linkqueue)
	t.Log(node)
	node = lq.Pop(linkqueue)
	t.Log(node)
	node = lq.Pop(linkqueue)
	t.Log(node)
	node = lq.Pop(linkqueue)
	t.Log(node)
	node = lq.Pop(linkqueue)
	t.Log(node)
	_ = lq.Push(linkqueue, 4)
	_ = lq.Push(linkqueue, 5)
	_ = lq.Push(linkqueue, 6)
	node = lq.Pop(linkqueue)
	t.Log(node)
	lq.Print(linkqueue)
}
