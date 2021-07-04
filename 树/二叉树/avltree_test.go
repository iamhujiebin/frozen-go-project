package 二叉树

import (
	"fmt"
	"testing"
)

var avl AvlBinarySortTree
var avlTree *bNode

func init() {
	avlTree = avl.InitTree(8, 2, 1, 6, 4, 9, 5, 0, 3, 7)
}

func TestAvlBinarySortTree_Insert(t *testing.T) {
	res := make([]interface{}, 0)
	avl.InOrder(avlTree, &res)
	fmt.Println(res)
}

func TestAvlBinarySortTree_Delete(t *testing.T) {
	res := make([]interface{}, 0)
	avl.InOrder(avlTree, &res)
	fmt.Println(res)

	avl.Delete(avlTree, 5) // 用bst的删除节点方法会导致不平衡
	res = make([]interface{}, 0)
	avl.InOrder(avlTree, &res)
	fmt.Println(res)
}
