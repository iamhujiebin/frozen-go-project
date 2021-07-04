package 二叉树

// 平衡二叉排序树
// 两个大佬的英文缩写AVL
// 是搜索性能最好的二叉排序树
// 定义:任意节点的左右子树的高度差不大于1的二叉排序树
// 所以bNode多一个height的参数
type AvlBinarySortTree struct {
	BinarySortTree
}

func (a AvlBinarySortTree) InitTree(datas ...interface{}) *bNode {
	var root *bNode
	for _, data := range datas {
		root = a.Insert(root, data)
	}
	return root
}

// 插入操作
// 口诀:
// 1.左左:向右旋转
// 2.左右:左右旋转
// 3.右右:向左旋转
// 4.右左:右左旋转
// ps:旋转后记得重新计算高度
func (a AvlBinarySortTree) Insert(root *bNode, data interface{}) *bNode {
	if root == nil {
		return &bNode{
			data:   data,
			left:   nil,
			right:  nil,
			height: 1,
		}
	}
	if data.(int) > root.ToInt() {
		root.right = a.Insert(root.right, data)
	} else {
		root.left = a.Insert(root.left, data)
	}
	root.height = max(a.getHeight(root.left), a.getHeight(root.right)) + 1
	leftHeight := a.getHeight(root.left)
	rightHeight := a.getHeight(root.right)
	if leftHeight-rightHeight == 2 { // 左子树不平衡
		if data.(int) < root.left.ToInt() { // 左左
			root = a.rotateRight(root)
		} else { // 左右
			root = a.rotateLeftRight(root)
		}
	}
	if leftHeight-rightHeight == -2 { // 右子树不平衡
		if data.(int) < root.right.ToInt() { // 右左
			root = a.rotateRightLeft(root)
		} else { // 右右
			root = a.rotateLeft(root)
		}
	}
	return root
}

func (a AvlBinarySortTree) getHeight(node *bNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 左旋转
//    9                   10
//     \                 /  \
//      10   ----->     9   11
//      / \              \
//    10.5  11           10.5
// 注意需要更新高度
func (a AvlBinarySortTree) rotateLeft(root *bNode) *bNode {
	if root == nil || root.right == nil {
		return nil
	}
	right := root.right
	root.right = right.left
	right.left = root
	root.height = max(a.getHeight(root.left), a.getHeight(root.right)) + 1
	right.height = max(a.getHeight(right.left), a.getHeight(right.right)) + 1
	return right
}

// 右旋转
//     9                8
//    /                / \
//   8      ----->    7   9
//  / \                   /
// 7  8.5                8.5
// 注意需要更新高度
func (a AvlBinarySortTree) rotateRight(root *bNode) *bNode {
	if root == nil || root.left == nil {
		return nil
	}
	left := root.left
	root.left = left.right
	left.right = root
	root.height = max(a.getHeight(root.left), a.getHeight(root.right)) + 1
	left.height = max(a.getHeight(left.left), a.getHeight(left.right)) + 1
	return left
}

// 左右旋转
func (a AvlBinarySortTree) rotateLeftRight(root *bNode) *bNode {
	if root == nil || root.left == nil {
		return nil
	}
	root.left = a.rotateLeft(root.left)
	return a.rotateRight(root)
}

// 右左旋转
func (a AvlBinarySortTree) rotateRightLeft(root *bNode) *bNode {
	if root == nil || root.right == nil {
		return nil
	}
	root.right = a.rotateRight(root.right)
	return a.rotateLeft(root)
}
