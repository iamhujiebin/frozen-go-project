package 二叉树

import . "frozen-go-project/utils/queuestack"

// 二叉排序树
// 定义:左子树每一个节点都比根节点小,右子树每一个节点都比根节点大
type BinarySortTree struct {
}

func (b BinarySortTree) InitTree(datas ...interface{}) *bNode {
	root := &bNode{
		data:  nil,
		left:  nil,
		right: nil,
	}
	for _, data := range datas {
		b.Insert(root, data)
	}
	return root
}

// 比当前节点大就右移动|作为右子节点
// 比当前节点小就左移动|作为左子节点
func (b BinarySortTree) Insert(root *bNode, data interface{}) {
	if root == nil {
		return
	}
	// 空树
	if root.data == nil && root.left == nil && root.right == nil {
		root.data = data
		return
	}
	cur := root
	for {
		if data.(int) >= cur.ToInt() {
			if cur.right == nil {
				cur.right = &bNode{data: data}
				return
			}
			cur = cur.right
		} else {
			if cur.left == nil {
				cur.left = &bNode{data: data}
				return
			}
			cur = cur.left
		}
	}
}

// 如果删除根节点,直接删除
func (b BinarySortTree) Delete(root *bNode, data interface{}) {
	panic("implement me")
}

func (b BinarySortTree) LevelOrder(root *bNode) (res []interface{}) {
	if root == nil {
		return nil
	}
	lq := new(LinkQueue)
	linkQueue, _ := lq.Init(0)
	_ = lq.Push(linkQueue, root)
	for !lq.IsEmpty(linkQueue) {
		n := lq.Pop(linkQueue).(*bNode)
		res = append(res, n.data)
		if n.left != nil {
			_ = lq.Push(linkQueue, n.left)
		}
		if n.right != nil {
			_ = lq.Push(linkQueue, n.right)
		}
	}
	return
}

func (b BinarySortTree) PreOrder(root *bNode, i *[]interface{}) {
	if root == nil {
		return
	}
	*i = append(*i, root.data)
	b.PreOrder(root.left, i)
	b.PreOrder(root.right, i)
}

func (b BinarySortTree) InOrder(root *bNode, i *[]interface{}) {
	if root == nil {
		return
	}
	b.InOrder(root.left, i)
	*i = append(*i, root.data)
	b.InOrder(root.right, i)
}

func (b BinarySortTree) PostOrder(root *bNode, i *[]interface{}) {
	if root == nil {
		return
	}
	b.PostOrder(root.left, i)
	b.PostOrder(root.right, i)
	*i = append(*i, root.data)
}

func (b BinarySortTree) TreeDepth(root *bNode) int {
	if root == nil {
		return 0
	}
	left := b.TreeDepth(root.left)
	right := b.TreeDepth(root.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func (b BinarySortTree) RestoreBTree(i []interface{}, i2 []interface{}) *bNode {
	panic("implement me")
}
