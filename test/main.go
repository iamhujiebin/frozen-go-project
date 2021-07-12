package main

import (
	"fmt"
	ConcurrentSkipList "frozen-go-project/test/SkipList"
	"github.com/pyihe/go-skipList"
)

func main() {
	skipList, err := ConcurrentSkipList.NewConcurrentSkipList(4)
	if err != nil {
		fmt.Println(err)
	}

	// Insert index and value. The index must uint64 and value is interface.
	skipList.Insert(uint64(11), 11)
	skipList.Insert(uint64(13), 13)
	skipList.Insert(uint64(14), 14)
	skipList.Insert(uint64(10), 10)
	skipList.Insert(uint64(15), 15)
	skipList.Insert(uint64(17), 17)
	skipList.Insert(uint64(18), 18)
	skipList.Insert(uint64(16), 16)
	skipList.Insert(uint64(19), 19)
	skipList.Insert(uint64(20), 20)
	skipList.Insert(uint64(12), 12)

	// Search in skip list.
	if node, ok := skipList.Search(uint64(1)); ok {
		fmt.Printf("index:%v value:%v\\n", node.Index(), node.Value())
	}

	// Delete by index.
	skipList.Delete(uint64(2))

	// Get the level of skip list.
	_ = skipList.Level()

	// Get the length of skip list.
	_ = skipList.Length()

	// Iterate each node in skip list.
	skipList.ForEach(func(node *ConcurrentSkipList.Node) bool {
		fmt.Printf("index:%v value:%v\\n", node.Index(), node.Value())
		return true
	})

	// Select top 10 nodes of skip list.
	nodes := skipList.Sub(0, 10)
	fmt.Printf("nodes:%+v\n", nodes)
}

func zset() {

	ss := go_skipList.NewSkipList()
	ss.InsertByEle("k1", 10, nil)
	_, _ = ss.InsertByEleArray("k2", 10.1, "this is k2", "k3", 1.1, nil)
	nodes, err := ss.GetElementByRank(1)
	if err != nil {
		//handle err
	}

	//output: mem: k2, score: 10.1, data: this is k2
	fmt.Printf("mem: %s, score: %v, data: %v\n", nodes[0].Name(), nodes[0].Score(), nodes[0].Data())
}
