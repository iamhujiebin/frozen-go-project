package queuestack

import (
	"fmt"
	"sync"
)

// 顺序栈-用数组实现栈
// 线程安全:加个锁
type SeqStack struct {
	lock sync.RWMutex // 加了锁之后,函数需要用pointer,因为golang值传递
}

func (s *SeqStack) Init(cap int, nodes ...*Node) (*QueueStack, error) {
	stack := &QueueStack{
		Len:      0,
		Cap:      cap,
		arr:      make([]*Node, cap),
		topIndex: -1,
	}
	for _, n := range nodes {
		err := s.Push(stack, n)
		if err != nil {
			return nil, err
		}
	}
	return stack, nil
}

func (s *SeqStack) Destroy(stack *QueueStack) {
	stack = nil
}

func (s *SeqStack) Clear(stack *QueueStack) {
	stack, _ = s.Init(stack.Cap)
}

func (s *SeqStack) Push(stack *QueueStack, node *Node) error {
	if stack == nil {
		return ErrNoInit
	}
	if node == nil {
		return ErrNode
	}
	if s.IsFull(stack) {
		return ErrFull
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	stack.topIndex++
	stack.Len++
	stack.arr[stack.topIndex] = node
	return nil
}

func (s *SeqStack) Pop(stack *QueueStack) *Node {
	if stack == nil {
		return nil
	}
	if s.IsEmpty(stack) {
		return nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	stack.Len--
	node := stack.arr[stack.topIndex]
	stack.arr[stack.topIndex] = nil // 需要置空,print会用
	stack.topIndex--
	return node
}

func (s *SeqStack) IsEmpty(stack *QueueStack) bool {
	// 读锁
	s.lock.RLock()
	defer s.lock.RUnlock()
	return stack.Len == 0
}

func (s *SeqStack) IsFull(stack *QueueStack) bool {
	if stack == nil {
		return false
	}
	s.lock.RLock()
	s.lock.RUnlock()
	return stack.Len >= stack.Cap
}

func (s *SeqStack) Length(stack *QueueStack) int {
	if stack == nil {
		return 0
	}
	return stack.Len
}

func (s *SeqStack) GetTop(stack *QueueStack) *Node {
	if stack == nil {
		return nil
	}
	s.lock.RLock()
	s.lock.RUnlock()
	return stack.arr[stack.topIndex]
}

func (s *SeqStack) Print(stack *QueueStack) {
	if stack == nil {
		return
	}
	var datas []interface{}
	for i := 0; i < stack.Len; i++ {
		datas = append(datas, stack.arr[i].Data)
	}
	fmt.Printf("Len:%v,Cap:%v,topIndex:%v,datas:%v",
		stack.Len, stack.Cap, stack.topIndex, datas)
}
