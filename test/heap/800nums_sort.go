package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type RowColVal struct {
	Row int
	Col int
	Val int
}

type IntHeap []RowColVal

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(RowColVal))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	test2()
	test1()
}

func test2() {
	start := time.Now()
	defer func() {
		fmt.Printf("sort cost:%v\n", time.Now().Sub(start).Seconds())
	}()
	var num []int
	for i := 0; i < 8000000; i++ {
		num = append(num, rand.Intn(10000000))
	}
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
}

func test1() {
	start := time.Now()
	defer func() {
		fmt.Printf("merge cost:%v", time.Now().Sub(start).Seconds())
	}()
	var num [8000000]int
	for i := 0; i < 8000000; i++ {
		num[i] = rand.Intn(10000000)
	}
	tNum := 8
	count := 100000
	var an [][]int
	for i := 0; i < tNum; i++ {
		an = append(an, num[i*count:(i+1)*count])
	}
	var wg sync.WaitGroup
	wg.Add(len(an))
	for k := range an {
		k := k
		go func() {
			defer wg.Done()
			sort.Slice(an[k], func(i, j int) bool {
				return an[k][i] < an[k][j]
			})
		}()
	}
	wg.Wait()
	var n int
	for _, v := range an {
		n += len(v)
	}
	intHeap := new(IntHeap)
	heap.Init(intHeap)
	for i, v := range an {
		heap.Push(intHeap, RowColVal{
			Row: i,
			Col: 0,
			Val: v[0],
		})
	}
	// 堆排序合并多个有序数组
	var merge = make([]int, n)
	for idx := range merge {
		min := heap.Pop(intHeap).(RowColVal)
		// 总数组加数据
		merge[idx] = min.Val
		// 子数组偏移
		if min.Col+1 < len(an[min.Row]) {
			heap.Push(intHeap, RowColVal{
				Row: min.Row,
				Col: min.Col + 1,
				Val: an[min.Row][min.Col+1],
			})
		}
	}
}
