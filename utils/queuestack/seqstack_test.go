package queuestack

import "testing"

var sStack *QueueStack
var ss SeqStack

func init() {
	sStack, _ = ss.Init(10, 1, 5, 9)
}

func TestSeqStack_Init(t *testing.T) {
	ss.Print(sStack)
}

func TestSeqStack_Clear(t *testing.T) {
	ss.Clear(sStack)
}

func TestSeqStack_Destroy(t *testing.T) {
	ss.Destroy(sStack)
}

func TestSeqStack_GetTop(t *testing.T) {
	t.Log(ss.GetTop(sStack))
	ss.Print(sStack)
}

func TestSeqStack_IsEmpty(t *testing.T) {
	t.Log(ss.IsEmpty(sStack))
}

func TestSeqStack_IsFull(t *testing.T) {
	t.Log(ss.IsFull(sStack))
}

func TestSeqStack_Length(t *testing.T) {
	t.Log(ss.Length(sStack))
}

func TestSeqStack_Pop(t *testing.T) {
	t.Log(ss.Pop(sStack))
	ss.Print(sStack)
}

func TestSeqStack_Push(t *testing.T) {
	err := ss.Push(sStack, 99)
	t.Log(err)
	ss.Print(sStack)
}
