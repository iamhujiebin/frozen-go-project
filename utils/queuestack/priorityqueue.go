package queuestack

import "sync"

// 优先队列
// 也是队列,有基本的Push/Pop
// Head指向优先级最高的节点
type PriorityQueue struct {
	lock sync.Mutex
}

func (p *PriorityQueue) Init() *QueueStack {
	return &QueueStack{
		Len:  0,
		head: nil,
	}
}

func (p *PriorityQueue) Push(queue *QueueStack, priority int, data interface{}) {
	if queue == nil {
		return
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	queue.Len++
	if queue.head == nil {
		queue.head = &node{
			Data:     data,
			Next:     nil,
			priority: priority,
		}
		return
	}
	node := &node{
		Data:     data,
		Next:     nil,
		priority: priority,
	}
	// 头节点
	if node.priority > queue.head.priority {
		node.Next = queue.head
		queue.head = node
		return
	}
	cur := queue.head
	for cur.Next != nil {
		if priority > cur.Next.priority {
			node.Next = cur.Next
			cur.Next = node
			return
		}
		cur = cur.Next
	}
	// 来到这,就是node的priority最小
	cur.Next = node
}

func (p *PriorityQueue) Pop(queue *QueueStack) interface{} {
	if queue == nil || queue.head == nil {
		return nil
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	queue.Len--
	n := queue.head
	queue.head = queue.head.Next
	return n.Data
}

func (p *PriorityQueue) IsEmpty(queue *QueueStack) bool {
	if queue == nil {
		return true
	}
	return queue.head == nil
}

func (p *PriorityQueue) Length(queue *QueueStack) int {
	if queue == nil {
		return 0
	}
	return queue.Len
}

func (p *PriorityQueue) GetTop(queue *QueueStack) interface{} {
	if queue == nil || queue.head == nil {
		return nil
	}
	return queue.head.Data
}
