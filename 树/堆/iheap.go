package 堆

type Element interface {
	Priority() int
	Bigger(Element) bool
}

// 堆:完全二叉树
// 可以用数组来表示树
// 公式:
// 1. son=dad*2+1 son2=dad*2+2
// 2. dad=(son-1)/2
// 操作:
// 1. 上浮 heapifyUp
// 2. 下沉 heapifyDown
type IHeap interface {
	InitHeap(cap int) // 初始化堆
	Push(Element)     // 加元素
	Pop() Element     // 删堆顶
	GetTop() Element  // 获取堆顶
	heapifyUp()       // 堆化-上浮
	heapifyDown()     // 堆化-下沉
}
