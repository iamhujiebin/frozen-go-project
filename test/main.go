package main

import (
	"fmt"
	"frozen-go-project/树/二叉树"
)

var bst 二叉树.BinarySortTree

func main() {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	println(INT_MAX, INT_MIN)

	tree := bst.InitTree(6, 2, 5, 4, 3, 4, 5, 7, 0, 2, 3)
	fmt.Println("tree depth:", bst.TreeDepth(tree))
	res := make([]interface{}, 0)
	bst.InOrder(tree, &res)
	fmt.Printf("inOrder:%v\n", res)
}
