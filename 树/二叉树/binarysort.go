package 二叉树

import . "frozen-go-project/utils/queuestack"

// 二叉排序树
// 定义:左子树每一个节点都比根节点小,右子树每一个节点都比根节点大
type BinarySortTree struct {
}

func (b BinarySortTree) InitTree(datas ...interface{}) *BNode {
	root := &BNode{
		Data:  nil,
		Left:  nil,
		Right: nil,
	}
	for _, node := range datas {
		b.Insert(root, &BNode{Data: node})
	}
	return root
}

// 比当前节点大就右移动|作为右子节点
// 比当前节点小就左移动|作为左子节点
func (b BinarySortTree) Insert(root *BNode, node *BNode) {
	if root == nil || node == nil {
		return
	}
	// 空树
	if root.Data == nil && root.Left == nil && root.Right == nil {
		root.Data = node.Data
		return
	}
	cur := root
	for {
		if node.ToInt() >= cur.ToInt() {
			if cur.Right == nil {
				cur.Right = node
				return
			}
			cur = cur.Right
		} else {
			if cur.Left == nil {
				cur.Left = node
				return
			}
			cur = cur.Left
		}
	}
}

// 如果删除根节点,直接删除
func (b BinarySortTree) Delete(root *BNode, node *BNode) {
	panic("implement me")
}

func (b BinarySortTree) LevelOrder(root *BNode) (res []interface{}) {
	if root == nil {
		return nil
	}
	lq := new(LinkQueue)
	linkQueue, _ := lq.Init(0)
	_ = lq.Push(linkQueue, &Node{Data: root})
	for !lq.IsEmpty(linkQueue) {
		node := lq.Pop(linkQueue).Data.(*BNode)
		res = append(res, node.Data)
		if node.Left != nil {
			_ = lq.Push(linkQueue, &Node{Data: node.Left})
		}
		if node.Right != nil {
			_ = lq.Push(linkQueue, &Node{Data: node.Right})
		}
	}
	return
}

func (b BinarySortTree) PreOrder(root *BNode, i *[]interface{}) {
	if root == nil {
		return
	}
	*i = append(*i, root.Data)
	b.PreOrder(root.Left, i)
	b.PreOrder(root.Right, i)
}

func (b BinarySortTree) InOrder(root *BNode, i *[]interface{}) {
	if root == nil {
		return
	}
	b.InOrder(root.Left, i)
	*i = append(*i, root.Data)
	b.InOrder(root.Right, i)
}

func (b BinarySortTree) PostOrder(root *BNode, i *[]interface{}) {
	if root == nil {
		return
	}
	b.PostOrder(root.Left, i)
	b.PostOrder(root.Right, i)
	*i = append(*i, root.Data)
}

func (b BinarySortTree) TreeDepth(root *BNode) int {
	if root == nil {
		return 0
	}
	left := b.TreeDepth(root.Left)
	right := b.TreeDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func (b BinarySortTree) RestoreBTree(i []interface{}, i2 []interface{}) *BNode {
	panic("implement me")
}
