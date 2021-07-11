package tree

import "testing"

type ConcreteElement struct {
	val int
}

func (c ConcreteElement) Value() interface{} {
	return c.val
}

func (c ConcreteElement) Compare(element Element) int {
	if c.val == element.Value().(int) {
		return 0
	}
	if c.val > element.Value().(int) {
		return 1
	} else {
		return -1
	}
}

func TestNewAVLTree(t *testing.T) {
	avl := NewAVLTree(ConcreteElement{val: 1})
	avl, _ = avl.AddNode(ConcreteElement{val: 2})
	avl, _ = avl.AddNode(ConcreteElement{val: 3})
	avl, _ = avl.AddNode(ConcreteElement{val: 4})
	avl, _ = avl.AddNode(ConcreteElement{val: 5})
	avl, _ = avl.AddNode(ConcreteElement{val: 6})
	avl, _ = avl.AddNode(ConcreteElement{val: 7})
	avl, _ = avl.AddNode(ConcreteElement{val: 8})
	avl, _ = avl.AddNode(ConcreteElement{val: 9})
	println(avl.String())
	avl, _ = avl.RemoveNode(ConcreteElement{val: 6})
	println(avl.String())
}

func TestNewRedBlackTree(t *testing.T) {
	rbTree := RedBlackTree{Root: nil}
	rbTree.AddNode(ConcreteElement{val: 1})
	rbTree.AddNode(ConcreteElement{val: 2})
	rbTree.AddNode(ConcreteElement{val: 3})
	rbTree.AddNode(ConcreteElement{val: 4})
	rbTree.AddNode(ConcreteElement{val: 5})
	rbTree.AddNode(ConcreteElement{val: 6})
	rbTree.AddNode(ConcreteElement{val: 7})
	rbTree.AddNode(ConcreteElement{val: 8})
	rbTree.AddNode(ConcreteElement{val: 9})
	println(rbTree.String())
	rbTree.RemoveNode(ConcreteElement{val: 7})
	println(rbTree.String())
	rbTree.RemoveNode(ConcreteElement{val: 3})
	println(rbTree.String())
}
