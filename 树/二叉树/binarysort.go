package 二叉树

import (
	. "frozen-go-project/utils/queuestack"
)

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

// 普通的查找就是递归查找O(N)
// 因为是二叉排序树,可以做到O(logN)
func (b BinarySortTree) Find(root *bNode, data interface{}) (n *bNode, pre *bNode) {
	if root == nil {
		return nil, nil
	}
	cur := root
	for cur != nil {
		if cur.data == data {
			n = cur
			return
		}
		if data.(int) > cur.ToInt() {
			pre = cur
			cur = cur.right
		} else {
			pre = cur
			cur = cur.left
		}
	}
	return
}

// 因为可能删除root节点,golang值传递(没有**指针的指针的操作),所以只能返回一个newRoot
// 先查找对应的节点位置,和前驱节点
// 如果前驱节点为空:
// 		即要删除根节点
// 如果前驱节点不是空:
//		即要删除非根节点
// 对于删除操作:
//		根节点:newRoot指向新的根
//		非根节点:前驱节点的左或右节点指向新的节点
// 记住口诀:
// 1.如果删除节点是叶子节点,直接删除
// 2.如果删除节点只有左/右子树,则直接用左/右子树替换自己
// 3.如果删除的节点有左右子树,那就用右子树的最左的节点(因为右子树,最小的值是最左的节点)替换自己
// 	3.1 注意此处最左的节点可能会有右子树,替换时候也是需要记录左子树的前驱节点(ppre)去保留右子树
// 上面的口诀,要运用在删除根节点和非根节点上
// ps:重点对象:前驱节点(pre),目标删除节点(n),以及两者的左右关系(leftNode)
func (b BinarySortTree) Delete(root *bNode, data interface{}) (newRoot *bNode) {
	if root == nil {
		return
	}
	newRoot = root
	n, pre := b.Find(root, data)
	if n == nil {
		return
	}
	// n!=nil 节点找到了

	// 没有前驱节点,说明是要删除根节点
	if pre == nil {
		// 只有根节点
		if root.left == nil && root.right == nil {
			root.data = nil
			return
		}
		// 只有左|右子树
		if root.left != nil && root.right == nil {
			newRoot = root.left
			return
		}
		if root.right != nil && root.left == nil {
			newRoot = root.right
			return
		}
		// 左右子树都有,找右子树的最左节点作为根节点
		cur := root.right
		var ppre *bNode
		for cur.left != nil {
			ppre = cur
			cur = cur.left
		}
		// 此时cur是右子树最左侧节点
		// cur肯定没有左节点,但是可能右节点
		if ppre != nil {
			// cur右节点归到ppre的左节点
			ppre.left = cur.right
			cur.left = root.left
			cur.right = root.right
		} else {
			// ppre为空,就是root.right没有左子树了
			cur.left = root.left
		}
		newRoot = cur
		return
	}
	// pre不是nil,就是要删除非根节点,root不用变
	// 目标节点是左节点还是右节点
	leftNode := true
	if pre.right != nil && pre.right.data == data {
		leftNode = false
	}
	// 叶子节点
	if n.left == nil && n.right == nil {
		if leftNode {
			pre.left = nil
		} else {
			pre.right = nil
		}
		return
	}
	// 删除节点只有左子树
	if n.left != nil && n.right == nil {
		if leftNode {
			pre.left = n.left
		} else {
			pre.right = n.left
		}
		return
	}
	// 删除节点只有右子树
	if n.right != nil && n.left == nil {
		if leftNode {
			pre.left = n.right
		} else {
			pre.right = n.right
		}
		return
	}
	// 有左右子树,那就用右子树的最左侧节点
	cur := n.right
	var ppre *bNode
	for cur.left != nil {
		ppre = cur
		cur = cur.left
	}
	// 来到这,cur是最左侧的节点,ppre是最左侧节点/空节点
	// cur要换掉被删除的节点
	if ppre != nil {
		ppre.left = cur.right
		cur.left = n.left
		cur.right = n.right
	} else {
		cur.left = n.left
	}
	if leftNode {
		pre.left = cur
	} else {
		pre.right = cur
	}
	return
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
