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

func (l *LinkQueue) Init(cap int, datas ...interface{}) (*QueueStack, error) {
	queue := &QueueStack{
		Len:  0,
		head: nil,
		tail: nil,
	}
	for _, node := range datas {
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

func (l *LinkQueue) Push(queue *QueueStack, data interface{}) error {
	if queue == nil {
		return ErrNoInit
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	n := &node{Data: data}
	if queue.Len == 0 {
		queue.head, queue.tail = n, n
	} else {
		queue.tail.Next = n
		queue.tail = n
	}
	queue.Len++
	return nil
}

func (l *LinkQueue) Pop(queue *QueueStack) interface{} {
	if queue == nil {
		return nil
	}
	if l.IsEmpty(queue) {
		return nil
	}
	queue.Len--
	n := queue.head
	queue.head = n.Next
	if queue.Len == 0 {
		queue.tail = nil
	}
	return n.Data
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

func (l *LinkQueue) GetTop(queue *QueueStack) interface{} {
	if queue == nil || queue.head == nil {
		return nil
	}
	return queue.head.Data
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
