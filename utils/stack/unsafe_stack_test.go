package stack

import "testing"

var stack *Stack
var u UnSafeStack

func init() {
	stack, _ = u.InitStack(10, &Node{Data: 1}, &Node{Data: 5}, &Node{Data: 9})
}

func TestUnSafeStack_InitStack(t *testing.T) {
	u.PrintStack(stack)
}

func TestUnSafeStack_ClearStack(t *testing.T) {
	u.ClearStack(stack)
}

func TestUnSafeStack_GetTop(t *testing.T) {
	t.Log(u.GetTop(stack))
}

func TestUnSafeStack_IsEmpty(t *testing.T) {
	t.Log(u.IsEmpty(stack))
}

func TestUnSafeStack_IsFull(t *testing.T) {
	t.Log(u.IsFull(stack))
}

func TestUnSafeStack_DestroyStack(t *testing.T) {
	u.DestroyStack(stack)
}

func TestUnSafeStack_Length(t *testing.T) {
	t.Log(u.Length(stack))
}

func TestUnSafeStack_Pop(t *testing.T) {
	u.PrintStack(stack)
	t.Log(u.Pop(stack))
	u.PrintStack(stack)
}

func TestUnSafeStack_Push(t *testing.T) {
	u.PrintStack(stack)
	_ = u.Push(stack, &Node{Data: 99})
	u.PrintStack(stack)
}
