package queuestack

import (
	"fmt"
	"testing"
)

var (
	pq            = PriorityQueue{}
	priorityQueue = pq.Init()
)

func TestPriorityQueue_Init(t *testing.T) {
	pq.Push(priorityQueue, 10, 10)
	pq.Push(priorityQueue, 9, 9)
	pq.Push(priorityQueue, 8, 8)
	pq.Push(priorityQueue, 7, 7)
	pq.Push(priorityQueue, 18, 18)
	pq.Push(priorityQueue, 20, 20)
	pq.Push(priorityQueue, 11, 11)
	pq.Push(priorityQueue, 13, 13)
	fmt.Println(pq.Pop(priorityQueue))
	fmt.Println(pq.Pop(priorityQueue))
	fmt.Println(pq.Pop(priorityQueue))
	fmt.Println(pq.Pop(priorityQueue))
	fmt.Println(pq.Pop(priorityQueue))
}
