package queuestack

import (
	"fmt"
	"sync"
)

// 链栈-用链表实现栈
type LinkStack struct {
	lock sync.RWMutex
}

func (u *LinkStack) Print(stack *QueueStack) {
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
		stack.Len, stack.Cap, stack.top, datas)
}

func (u *LinkStack) Init(cap int, datas ...interface{}) (*QueueStack, error) {
	s := &QueueStack{
		top: nil,
		Len: 0,
		Cap: cap,
	}
	for _, data := range datas {
		err := u.Push(s, data)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (u *LinkStack) Destroy(stack *QueueStack) {
	stack = nil
}

func (u *LinkStack) Clear(stack *QueueStack) {
	stack.top = nil
	stack.Len = 0
}

func (u *LinkStack) Push(stack *QueueStack, data interface{}) error {
	if stack == nil {
		return ErrNoInit
	}
	if u.IsFull(stack) {
		return ErrFull
	}
	n := &node{Data: data}
	u.lock.Lock()
	u.lock.Unlock()
	stack.Len++
	n.Next = stack.top
	stack.top = n
	return nil
}

func (u *LinkStack) Pop(stack *QueueStack) interface{} {
	if stack == nil {
		return nil
	}
	if u.IsEmpty(stack) {
		return nil
	}
	u.lock.Lock()
	u.lock.Unlock()
	stack.Len--
	n := stack.top
	stack.top = n.Next
	return n.Data
}

func (u *LinkStack) IsEmpty(stack *QueueStack) bool {
	if stack == nil {
		return true
	}
	u.lock.RLock()
	u.lock.RUnlock()
	return stack.Len <= 0
}

func (u *LinkStack) IsFull(stack *QueueStack) bool {
	// 链栈不需要满
	return false
}

func (u *LinkStack) Length(stack *QueueStack) int {
	if stack == nil {
		return 0
	}
	u.lock.RLock()
	u.lock.RUnlock()
	return stack.Len
}

func (u *LinkStack) GetTop(stack *QueueStack) interface{} {
	if stack == nil || stack.top == nil {
		return nil
	}
	u.lock.RLock()
	u.lock.RUnlock()
	return stack.top.Data
}
