package 链表

import "fmt"

type SingleLinkList struct {
}

func (s SingleLinkList) DestroyList(head Head) {
	head = nil
	// 其他节点等gc
}

func (s SingleLinkList) InitList(nodes ...*Node) Head {
	head := &Node{
		Data: nil,
		Next: nil,
	}
	for k := range nodes {
		_ = s.PushBack(head, nodes[k])
	}
	return head
}

func (s SingleLinkList) ClearList(head Head) {
	if head == nil {
		return
	}
	head.Next = nil
	// 其他node等待GC
}

func (s SingleLinkList) PrintList(head Head) {
	var datas []interface{}
	p := head
	for p != nil {
		if p.Data != nil {
			datas = append(datas, p.Data)
		}
		p = p.Next
	}
	fmt.Println(datas)
}

func (s SingleLinkList) InsertListPos(head Head, pos int, node *Node) error {
	if head == nil {
		return ErrNoHead
	}
	if pos < 1 {
		return ErrNoPos
	}
	// 找到pos-1的位置,需要考虑head
	find := pos - 1
	n := 0
	cur := head
	for cur != nil && n < find {
		cur = cur.Next
		n++
	}
	if n != find {
		return ErrNoPos
	}
	node.Next = cur.Next
	cur.Next = node
	return nil
}

func (s SingleLinkList) PushFront(head Head, node *Node) error {
	if head == nil {
		return ErrNoHead
	}
	//return s.InsertListPos(head, 1, node) // 最好别复用了.因为这里会改变node的Next
	node.Next = head.Next
	head.Next = node
	return nil
}

func (s SingleLinkList) PushBack(head Head, node *Node) error {
	if head == nil {
		return ErrNoHead
	}
	p := head
	for p != nil && p.Next != nil {
		p = p.Next
	}
	p.Next = node
	return nil
}

func (s SingleLinkList) DeleteNodePos(head Head, pos int) error {
	if head == nil {
		return ErrNoHead
	}
	// 找到pos-1位置
	find := pos - 1
	n := 0
	cur := head
	for cur != nil && n < find {
		cur = cur.Next
		n++
	}
	if n != find {
		return ErrNoPos
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	} else {
		// pos是最后一个节点
		cur.Next = nil
	}
	return nil
}

func (s SingleLinkList) PopFront(head Head) *Node {
	if head == nil {
		return nil
	}
	var n *Node
	if head.Next != nil {
		n = head.Next
		head.Next = head.Next.Next
	}
	return n
}

func (s SingleLinkList) PopBack(head Head) *Node {
	if head == nil {
		return nil
	}
	cur := head
	// 找倒数第二个,需要两次Next
	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next
	}
	n := cur.Next
	cur.Next = nil
	return n
}

func (s SingleLinkList) LengthList(head Head) int {
	if head == nil {
		return 0
	}
	n := 0
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	return n
}

func (s SingleLinkList) IsEmpty(head Head) bool {
	if head == nil {
		return true
	}
	if head.Next != nil {
		return false
	}
	return true
}

func (s SingleLinkList) LocateNodePos(head Head, pos int) *Node {
	if head == nil {
		return nil
	}
	n := 0
	cur := head
	for cur != nil && n < pos {
		cur = cur.Next
		n++
	}
	if n != pos {
		return nil
	}
	return cur
}

func (s SingleLinkList) LocateNodeElem(head Head, data interface{}) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		if cur.Data == data {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (s SingleLinkList) InsertNextNode(head Head, data interface{}, node *Node) error {
	if head == nil {
		return ErrNoHead
	}
	cur := head
	for cur != nil {
		if cur.Data == data {
			node.Next = cur.Next
			cur.Next = node
			return nil
		}
		cur = cur.Next
	}
	return ErrNoPos
}

func (s SingleLinkList) InsertPriorNode(head Head, data interface{}, node *Node) error {
	if head == nil {
		return ErrNoHead
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Data == data {
			node.Next = cur.Next
			cur.Next = node
			return nil
		}
		cur = cur.Next
	}
	return ErrNoPos
}

func (s SingleLinkList) DeleteNode(head Head, data interface{}) error {
	if head == nil {
		return ErrNoHead
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			continue // 全删掉就要continue!
			// break:只删掉第一个
		}
		cur = cur.Next
	}
	return nil
}

// 断掉head
// 再遍历无头链表
// 每个都执行pushFront即可
// ps:可复用pushFront,但记得要保留一下cur.Next
func (s SingleLinkList) ReverseList(head Head) {
	if head == nil {
		return
	}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		//_ = s.PushFront(head, cur)
		cur.Next = head.Next
		head.Next = cur
		cur = next
	}
	return
}

func (s SingleLinkList) ReversePartList(head Head, pos int) error {
	if head == nil {
		return ErrNoHead
	}
	// 找到前一个节点
	find := pos - 1
	cur := head
	n := 0
	for cur != nil && n < find {
		n++
		cur = cur.Next
	}
	if n != find {
		return ErrNoPos
	}
	partHead := cur
	cur = partHead.Next
	// 断掉head
	partHead.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = partHead.Next
		partHead.Next = cur
		cur = next
	}
	return nil
}

func (s SingleLinkList) MergeList(head1 Head, head2 Head) (head Head) {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head = s.InitList()
	cur1, cur2 := head1.Next, head2.Next // 需要用无头链表！
	for cur1 != nil && cur2 != nil {
		next1, next2 := cur1.Next, cur2.Next
		if cur1.Data.(int) < cur2.Data.(int) {
			cur1.Next = nil            // 需要把Next干掉
			_ = s.PushBack(head, cur1) // 这里其实不用担心cur1被改变,因为函数里面没有操作cur1.Next
			cur1 = next1
		} else {
			cur2.Next = nil            // 同上！
			_ = s.PushBack(head, cur2) // 区别于PushFront
			cur2 = next2
		}
	}
	for cur1 != nil {
		_ = s.PushBack(head, cur1)
		cur1 = cur1.Next
	}
	for cur2 != nil {
		_ = s.PushBack(head, cur2)
		cur2 = cur2.Next
	}
	return
}
