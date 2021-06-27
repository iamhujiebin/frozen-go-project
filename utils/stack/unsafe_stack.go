package stack

import "fmt"

type UnSafeStack struct {
}

func (u UnSafeStack) PrintStack(stack *Stack) {
	if stack == nil {
		return
	}
	var datas []interface{}
	cur := stack.top
	for cur != nil {
		if cur.Data != nil {
			datas = append(datas, cur.Data)
		}
		cur = cur.Next
	}
	fmt.Printf("len:%v,cap:%v,top:%v,data:%v\n",
		stack.Len, stack.Cap, stack.top.Data, datas)
}

func (u UnSafeStack) InitStack(cap int, nodes ...*Node) (*Stack, error) {
	s := &Stack{
		top: nil,
		Len: 0,
		Cap: cap,
	}
	for _, n := range nodes {
		err := u.Push(s, n)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (u UnSafeStack) DestroyStack(stack *Stack) {
	stack = nil
}

func (u UnSafeStack) ClearStack(stack *Stack) {
	stack.top = nil
	stack.Len = 0
}

func (u UnSafeStack) Push(stack *Stack, node *Node) error {
	if stack == nil {
		return ErrNoInit
	}
	if node == nil {
		return ErrNode
	}
	if u.IsFull(stack) {
		return ErrFull
	}
	stack.Len++
	node.Next = stack.top
	stack.top = node
	return nil
}

func (u UnSafeStack) Pop(stack *Stack) *Node {
	if stack == nil {
		return nil
	}
	if u.IsEmpty(stack) {
		return nil
	}
	stack.Len--
	node := stack.top
	stack.top = node.Next
	return node
}

func (u UnSafeStack) IsEmpty(stack *Stack) bool {
	if stack == nil {
		return true
	}
	return stack.Len <= 0
}

func (u UnSafeStack) IsFull(stack *Stack) bool {
	return stack.Len >= stack.Cap
}

func (u UnSafeStack) Length(stack *Stack) int {
	if stack == nil {
		return 0
	}
	return stack.Len
}

func (u UnSafeStack) GetTop(stack *Stack) *Node {
	if stack == nil {
		return nil
	}
	return stack.top
}
