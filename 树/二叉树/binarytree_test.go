package 二叉树

import (
	"fmt"
	"testing"
)

var binaryTree *bNode
var bt BinaryTree

func init() {
	binaryTree = bt.InitTree(nil)
}

func TestBinaryTree_TreeDepth(t *testing.T) {
	t.Log(bt.TreeDepth(binaryTree))
}

func TestBinaryTree_LevelOrder(t *testing.T) {
	fmt.Printf("levelOrder:%v\n", bt.LevelOrder(binaryTree))
}

func TestBinaryTree_PreOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bt.PreOrder(binaryTree, &res)
	fmt.Printf("PreOrder:%v\n", res)
	res = make([]interface{}, 0)
	bt.PreOrder1(binaryTree, &res)
	fmt.Printf("PreOrder1:%v\n", res)
}

func TestBinaryTree_InOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bt.InOrder(binaryTree, &res)
	fmt.Printf("InOrder:%v\n", res)
	res = make([]interface{}, 0)
	bt.InOrder1(binaryTree, &res)
	fmt.Printf("InOrder1:%v\n", res)
}

// ps:后序遍历的非递归算法,不要求掌握,流程控制太难
func TestBinaryTree_PostOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bt.PostOrder(binaryTree, &res)
	fmt.Printf("PostOrder:%v\n", res)
}
