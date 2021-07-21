package 多叉树

func FindLowParent(root *nNode, n1, n2 interface{}) *nNode {
	if root == nil {
		return nil
	}
	if n1 == n2 && root.Data == n1 {
		return nil
	}
	// 同样的节点,那就返回其中一个的父节点,有就有无就无
	if n1 == n2 && root.Data != n1 {
		return findParent(root, n1)
	}
	// n1!=n2,循环找爹
	var p1, p2 *nNode
	p1 = findParent(root, n1)
	p2 = findParent(root, n2)
	if p1 == p2 {
		return p1
	}
	for p1 != nil && p2 != nil {
		p1 = findParent(root, p1)
		p2 = findParent(root, p2)
	}
	if p1 == nil || p2 == nil {
		return nil
	}
	return p1
}

func findParent(root *nNode, n interface{}) *nNode {
	if root == nil {
		return nil
	}
	cur := root
	for _, child := range cur.Nodes {
		if child.Data == n {
			return cur
		}
	}
	// 到下层
	for i := range cur.Nodes {
		parent := findParent(cur.Nodes[i], n)
		if parent != nil {
			return parent
		}
	}
	return nil
}
