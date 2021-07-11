package 堆

import "sync"

type Heap struct {
	Cap int
	Len int
	arr []Element

	lock sync.Mutex
}

func (h *Heap) InitHeap(cap int) {
	h.Cap = cap
	h.Len = 0
	h.arr = make([]Element, cap)
}

func (h *Heap) Push(i Element) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len >= h.Cap {
		// double 扩容
		h.Cap *= 2
		arr := make([]Element, h.Cap)
		copy(arr, h.arr)
		h.arr = arr
	}
	h.arr[h.Len] = i
	h.Len++
	h.heapifyUp()
}

func (h *Heap) Pop() Element {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len <= 0 {
		return nil
	}
	n := h.arr[0]
	h.arr[0] = h.arr[h.Len-1] // 替换第一个
	h.arr[h.Len-1] = nil      // 删掉最后一个
	h.Len--
	if h.Len > 0 {
		h.heapifyDown()
	}
	return n
}

func (h *Heap) GetTop() Element {
	if h.Len <= 0 {
		return nil
	}
	return h.arr[0]
}

func (h *Heap) heapifyUp() {
	son := h.Len - 1
	dad := (son - 1) / 2
	for dad != son {
		if h.arr[son].Bigger(h.arr[dad]) {
			h.arr[son], h.arr[dad] = h.arr[dad], h.arr[son]
		}
		son = dad
		dad = (son - 1) / 2
	}
}

func (h *Heap) heapifyDown() {
	dad := 0
	son := 2*dad + 1
	for son < h.Len-1 {
		if son+1 < h.Len-1 && h.arr[son+1].Bigger(h.arr[son]) {
			son++
		}
		if h.arr[son].Bigger(h.arr[dad]) {
			h.arr[son], h.arr[dad] = h.arr[dad], h.arr[son]
		}
		dad = son
		son = 2*dad + 1
	}
}
