package 多叉树

func FindLowParent(root *nNode, n1, n2 interface{}) *nNode {
	// 顺序表存储
	var p1, p2 []*nNode
	for p := findParent(root, n1); p != nil; p = findParent(root, n1) {
		p1 = append(p1, p)
		n1 = p.Data
	}
	for p := findParent(root, n2); p != nil; p = findParent(root, n2) {
		p2 = append(p2, p)
		n2 = p.Data
	}
	if len(p2) <= 0 {
		return nil
	}
	m2 := make(map[*nNode]struct{}, len(p2))
	for k := range p2 {
		m2[p2[k]] = struct{}{}
	}
	for k := range p1 {
		if _, ok := m2[p1[k]]; ok {
			return p1[k]
		}
	}
	return nil
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
