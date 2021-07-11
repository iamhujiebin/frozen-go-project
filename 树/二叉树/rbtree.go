package 二叉树

// 红黑树,本质也是二叉排序树
// 定义:
//	1. 任意节点的子树黑色节点的高度一致
//  2. 根节点是黑色
//  3. 红色节点不能相邻
// 口诀: 黑根黑叶红不邻,同祖等高只数黑
// 算法复杂度:
// 	1. 查找时间复杂度O(logN)
//	2. 插入最多两次旋转,删除最多3次旋转 O(1) + O(logN)[查找的复杂度]
//  3. AVL树的删除需要旋转O(logN)次,时间复杂度是O(logN) + O(logN)[查找]
type RBTree struct {
	BinarySortTree
}

func (rb RBTree) InitTree(datas ...interface{}) *bNode {
	var root *bNode
	for _, data := range datas {
		root = rb.Insert(root, data)
	}
	return root
}

// 插入节点的"七种"情况
// 0.插入的节点初始颜色是红色,根节点是黑色
// 1.树为空:直接插入黑色根节点 1⃣️
// 2.树不空:
//  2.0 查找插入位置
//	2.1 父节点是黑色:直接插入红色节点(不影响树高和红不邻) 2⃣️
//  2.2 父节点是红色:先插入新节点,再需"自平衡"-->其实就是一个函数,for循环遍历5种情况
//		2.2.1 叔伯节点是红色:爷父叔变色(高度不变),向上传递 3⃣️
//		2.2.2 叔伯节点是黑色:
//			2.2.2.1 左右 爷父新成< 4⃣️
//			2.2.2.2 左左 爷父新成/ 5⃣️
//			2.2.2.3 右左 爷父新成> 6⃣️
//			2.2.2.4 右右 爷父新成\ 7⃣️
// 其中情况的处理方式
// 1⃣️、2⃣️:直接操作即可
// 3⃣️~7⃣️:先插入,再"自平衡",自平衡从新节点入手
// 3⃣️:爷父叔变色,继续循环处理爷节点.(重新进入"自平衡",变4⃣️)
// 4⃣️:以父节点中心左旋,变5⃣️
// 5⃣️:爷父变色,以爷节点中心右旋
// 6⃣️:以父节点为中心右旋,变7⃣️
// 7⃣️:爷父变色,以爷节点中心左旋
func (rb RBTree) Insert(root *bNode, data interface{}) *bNode {
	// 1⃣️
	if root == nil {
		return &bNode{
			data:    data,
			left:    nil,
			right:   nil,
			isBlack: true,
			parent:  nil,
		}
	}
	// 新插入的红色节点
	node := &bNode{
		data:    data,
		left:    nil,
		right:   nil,
		isBlack: false, // 红色
		parent:  nil,
	}
	// 查找插入位置
	pos := root
	var parent *bNode
	isLeft := true
	for pos != nil {
		parent = pos
		if data.(int) < pos.ToInt() {
			isLeft = true
			pos = pos.left
		} else {
			isLeft = false
			pos = pos.right
		}
	}
	// pos==nil就是找到了
	node.parent = parent
	if isLeft {
		parent.left = node
	} else {
		parent.right = node
	}
	// 父节点是黑色,直接返回 2⃣️
	if parent.isBlack {
		return root
	}
	// 父节点是红色,需要"自平衡"
	rb.insertBalance(node)
	return root
}

// 自平衡,3⃣️~7⃣️都处理了
func (rb RBTree) insertBalance(node *bNode) {
	panic(node)
}

// 左旋
//    parent             parent
//    /                    /
//    9(node)             10
//     \                 /  \
//      10   ----->     9   11
//      / \              \
//    10.5  11           10.5
// 因为多了parent节点,所以不需要返回值(区别与avl)
func (rb RBTree) rotateLeft(node *bNode) {
	if node == nil {
		return
	}
	parent := node.parent
	right := node.right

	// 左旋
	node.right = right.left
	right.left = node

	// 重新设置parent
	if right.left != nil {
		right.left.parent = node
	}
	node.parent = right
	if parent == nil {
		right.parent = nil
	} else {
		if parent.left == node {
			parent.left = right
		} else {
			parent.right = right
		}
		right.parent = parent
	}
}

// 右旋
//    parent           parent
//      /                /
//     9(node)          8
//    /                / \
//   8      ----->    7   9
//  / \                   /
// 7  8.5                8.5
// 因为多了parent节点,所以不需要返回值(区别与avl)
func (rb RBTree) rotateRight(node *bNode) {
	if node == nil {
		return
	}
	parent := node.parent
	left := node.left

	// 右旋
	node.left = left.right
	left.right = node

	// 重设parent
	if left.right != nil {
		left.right.parent = node
	}
	node.parent = left
	if parent == nil {
		left.parent = nil
	} else {
		left.parent = parent
		if parent.left == node {
			parent.left = left
		} else {
			parent.right = left
		}
	}
}

// 步骤
// 1. 查找位置
// 2. 找替代节点 五种情况
//	 2.1 删除节点有两个子节点:找后继节点作为删除节点,值替换(后继节点:大于删除节点的最小节点,即右子树最左侧,或父节点左侧),进入2.2/2.3
//   2.2 删除节点(后继节点)还有子节点,优先选左节点为替代节点(没有就用右节点)
//   2.3 删除节点没有子节点,无需替代了
// 3. 自平衡(删除自平衡) 八种情况!
//  3.1 删除节点是黑色才需要自平衡,若是红色,不影响高度
// 4. 真正删除
func (rb RBTree) Delete(root *bNode, data interface{}) *bNode {
	return nil
}
