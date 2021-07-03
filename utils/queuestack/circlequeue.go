package queuestack

import (
	"fmt"
	"sync"
)

// 顺序队列,用数组实现
// 加上头尾指针
// 就成了循环队列！ps:channel 就是用循环队列实现的
type CircleQueue struct {
	lock sync.RWMutex
}

func (c *CircleQueue) Init(cap int, datas ...interface{}) (*QueueStack, error) {
	q := &QueueStack{
		Cap:       cap,
		arr:       make([]interface{}, cap),
		headIndex: -1,
		tailIndex: -1,
	}
	for _, node := range datas {
		_ = c.Push(q, node)
	}
	return q, nil
}

func (c *CircleQueue) Destroy(queue *QueueStack) {
	queue = nil
}

func (c *CircleQueue) Clear(queue *QueueStack) {
	queue, _ = c.Init(queue.Cap)
}

func (c *CircleQueue) Push(queue *QueueStack, data interface{}) error {
	if queue == nil {
		return ErrNoInit
	}
	if c.IsFull(queue) {
		return ErrFull
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	// 队列为空的时候,head/tail同移
	// 注意不能用c.IsEmpty()，会死锁
	if queue.Len == 0 {
		queue.headIndex++
		if queue.headIndex >= queue.Cap {
			queue.headIndex = 0
		}
	}
	queue.Len++
	queue.tailIndex++
	if queue.tailIndex >= queue.Cap {
		queue.tailIndex = 0
	}
	queue.arr[queue.tailIndex] = data
	return nil
}

func (c *CircleQueue) Pop(queue *QueueStack) interface{} {
	if queue == nil {
		return nil
	}
	if c.IsEmpty(queue) {
		return nil
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	queue.Len--
	data := queue.arr[queue.headIndex]
	queue.headIndex++
	if queue.headIndex >= queue.Cap {
		queue.headIndex = 0
	}
	// 队列为空的时候,head/tail同移
	if queue.Len == 0 {
		queue.tailIndex++
		if queue.tailIndex >= queue.Cap {
			queue.tailIndex = 0
		}
	}
	return data
}

func (c *CircleQueue) IsEmpty(queue *QueueStack) bool {
	if queue == nil {
		return false
	}
	c.lock.RLock()
	defer c.lock.RUnlock()
	return queue.Len == 0
}

func (c *CircleQueue) IsFull(queue *QueueStack) bool {
	if queue == nil {
		return false
	}
	c.lock.RLock()
	defer c.lock.RUnlock()
	return queue.Len >= queue.Cap
}

func (c *CircleQueue) Length(queue *QueueStack) int {
	if queue == nil {
		return 0
	}
	c.lock.RLock()
	defer c.lock.RUnlock()
	return queue.Len
}

func (c *CircleQueue) GetTop(queue *QueueStack) interface{} {
	if queue == nil {
		return nil
	}
	return queue.arr[queue.headIndex]
}

func (c *CircleQueue) Print(queue *QueueStack) {
	if queue == nil {
		return
	}
	datas := make([]interface{}, queue.Cap)
	for i := 0; i < queue.Cap; i++ {
		if queue.arr[i] == nil {
			datas[i] = nil
		} else {
			datas[i] = queue.arr[i]
		}
	}
	fmt.Printf("Cap:%v,Len:%v,headIndex:%v,tailIndex:%v,datas:%v\n",
		queue.Cap, queue.Len, queue.headIndex, queue.tailIndex, datas)
}
