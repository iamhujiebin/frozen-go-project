package queuestack

import "testing"

var linkStack *QueueStack
var ls LinkStack

func init() {
	linkStack, _ = ls.Init(10, &Node{Data: 1}, &Node{Data: 5}, &Node{Data: 9})
}

func TestLinkStack_InitStack(t *testing.T) {
	ls.Print(linkStack)
}

func TestLinkStack_ClearStack(t *testing.T) {
	ls.Clear(linkStack)
}

func TestLinkStack_GetTop(t *testing.T) {
	t.Log(ls.GetTop(linkStack))
}

func TestLinkStack_IsEmpty(t *testing.T) {
	t.Log(ls.IsEmpty(linkStack))
}

func TestLinkStack_IsFull(t *testing.T) {
	t.Log(ls.IsFull(linkStack))
}

func TestLinkStack_DestroyStack(t *testing.T) {
	ls.Destroy(linkStack)
}

func TestLinkStack_Length(t *testing.T) {
	t.Log(ls.Length(linkStack))
}

func TestLinkStack_Pop(t *testing.T) {
	ls.Print(linkStack)
	t.Log(ls.Pop(linkStack))
	ls.Print(linkStack)
}

func TestLinkStack_Push(t *testing.T) {
	ls.Print(linkStack)
	_ = ls.Push(linkStack, &Node{Data: 99})
	ls.Print(linkStack)
}
