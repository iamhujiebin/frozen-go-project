package queuestack

import (
	"fmt"
	"sync"
)

// 链式队列
// FIFO:先进先出
type LinkQueue struct {
	lock sync.Mutex
}

func (l *LinkQueue) Init(cap int, nodes ...*Node) (*QueueStack, error) {
	queue := &QueueStack{
		Len:  0,
		head: nil,
		tail: nil,
	}
	for _, node := range nodes {
		_ = l.Push(queue, node)
	}
	return queue, nil
}

func (l *LinkQueue) Destroy(queue *QueueStack) {
	queue = nil
}

func (l *LinkQueue) Clear(queue *QueueStack) {
	if queue == nil {
		return
	}
	queue.head, queue.tail = nil, nil
}

func (l *LinkQueue) Push(queue *QueueStack, node *Node) error {
	if queue == nil {
		return ErrNoInit
	}
	if node == nil {
		return ErrNode
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	if queue.Len == 0 {
		queue.head, queue.tail = node, node
	} else {
		queue.tail.Next = node
		queue.tail = node
	}
	queue.Len++
	return nil
}

func (l *LinkQueue) Pop(queue *QueueStack) *Node {
	if queue == nil {
		return nil
	}
	if l.IsEmpty(queue) {
		return nil
	}
	queue.Len--
	node := queue.head
	queue.head = node.Next
	if queue.Len == 0 {
		queue.tail = nil
	}
	return node
}

func (l *LinkQueue) IsEmpty(queue *QueueStack) bool {
	if queue == nil {
		return false
	}
	return queue.Len == 0
}

// 链式队列不需要满
func (l *LinkQueue) IsFull(queue *QueueStack) bool {
	return false
}

func (l *LinkQueue) Length(queue *QueueStack) int {
	if queue == nil {
		return 0
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	return queue.Len
}

func (l *LinkQueue) GetTop(queue *QueueStack) *Node {
	if queue == nil {
		return nil
	}
	return queue.head
}

func (l *LinkQueue) Print(queue *QueueStack) {
	if queue == nil {
		fmt.Printf("nil queue")
		return
	}
	var data []interface{}
	cur := queue.head
	for cur != nil {
		data = append(data, cur.Data)
		cur = cur.Next
	}
	fmt.Printf("Len:%v,head:%v,tail:%v,data:%v", queue.Len, queue.head, queue.tail, data)
}
