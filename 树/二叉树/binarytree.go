package 二叉树

import (
	. "frozen-go-project/utils/queuestack"
)

type BinaryTree struct {
}

/*
	 1
   /  \
  2    3
 /  \   /
 4   5 6
/ \  /\
7  8 9 0
*/
func (B BinaryTree) InitTree(datas ...interface{}) *BNode {
	//if len(datas) <= 0 {
	//	return nil
	//}
	n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 :=
		&BNode{Data: 1}, &BNode{Data: 2}, &BNode{Data: 3}, &BNode{Data: 4}, &BNode{Data: 5},
		&BNode{Data: 6}, &BNode{Data: 7}, &BNode{Data: 8}, &BNode{Data: 9}, &BNode{Data: 0}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left = n6
	n4.Left, n4.Right = n7, n8
	n5.Left, n5.Right = n9, n10
	return n1
}

func (B BinaryTree) Insert(root *BNode, node *BNode) {
	panic("implement me")
}

func (B BinaryTree) Delete(root *BNode, node *BNode) {
	panic("implement me")
}

// 利用一个队列
// root先放入
// 出一个节点，再把它的左右节点放进去
// 循环直到队列为空
func (B BinaryTree) LevelOrder(root *BNode) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}
	var queue []*BNode
	// root先入队
	queue = append(queue, root)
	for len(queue) > 0 {
		// queue出队一个
		// 把节点的左右子节点再入队即可
		node := queue[0]
		// 干掉前一个
		queue = queue[1:]
		res = append(res, node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func (B BinaryTree) PreOrder(root *BNode, res *[]interface{}) {
	if root == nil {
		return
	}
	*res = append(*res, root.Data)
	B.PreOrder(root.Left, res)
	B.PreOrder(root.Right, res)
	return
}

// 用栈来实现递归: 递去=入栈 归来=出栈
// 前序遍历:节点入栈前访问
// 步骤:
// 1. 从根节点,沿左边节点,依次入栈(递去),直到左节点为空。(可以准备归来了)
// 2. 出栈 ,pop出来的有右节点,把右节点当成子树,回到步骤1;如果没有右节点,回到步骤2
// 3. 直至栈为空
func (B BinaryTree) PreOrder1(root *BNode, res *[]interface{}) {
	if root == nil || res == nil {
		return
	}
	ls := LinkStack{}
	stack, _ := ls.Init(0)
	// 步骤1
	cur := root
	for cur != nil {
		// 节点入栈前访问,下同
		*res = append(*res, cur.Data)
		_ = ls.Push(stack, &node{Data: cur})
		cur = cur.Left
	}
	for !ls.IsEmpty(stack) {
		node := ls.Pop(stack)
		// 步骤3
		if node.Data.(*BNode).Right != nil {
			cur := node.Data.(*BNode).Right
			for cur != nil {
				*res = append(*res, cur.Data)
				_ = ls.Push(stack, &node{Data: cur})
				cur = cur.Left
			}
		}
	}
}

func (B BinaryTree) InOrder(root *BNode, res *[]interface{}) {
	if root == nil {
		return
	}
	B.InOrder(root.Left, res)
	*res = append(*res, root.Data)
	B.InOrder(root.Right, res)
}

// 步骤同前序遍历
// 不同点: 出栈的时候访问数据
func (B BinaryTree) InOrder1(root *BNode, res *[]interface{}) {
	if root == nil || res == nil {
		return
	}
	ls := LinkStack{}
	stack, _ := ls.Init(0)
	// 步骤1
	cur := root
	for cur != nil {
		_ = ls.Push(stack, &node{Data: cur})
		cur = cur.Left
	}
	for !ls.IsEmpty(stack) {
		node := ls.Pop(stack)
		// 出栈的时候访问
		*res = append(*res, node.Data.(*BNode).Data)
		// 步骤3
		if node.Data.(*BNode).Right != nil {
			cur := node.Data.(*BNode).Right
			for cur != nil {
				_ = ls.Push(stack, &node{Data: cur})
				cur = cur.Left
			}
		}
	}
}

func (B BinaryTree) PostOrder(root *BNode, res *[]interface{}) {
	if root == nil {
		return
	}
	B.PostOrder(root.Left, res)
	B.PostOrder(root.Right, res)
	*res = append(*res, root.Data)
}

func (B BinaryTree) TreeDepth(root *BNode) int {
	if root == nil {
		return 0
	}
	leftDepth := B.TreeDepth(root.Left) + 1
	rightDepth := B.TreeDepth(root.Right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (B BinaryTree) RestoreBTree(i []interface{}, i2 []interface{}) *BNode {
	panic("implement me")
}
