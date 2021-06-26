package 链表

import "fmt"

/*
	没有带2的,空间复杂度O(n)
	有带2的,空间复杂度O(1)
*/

// 反向输出带头节点的单链表全部节点的值
func PrintReverse(head Head) {
	single := SingleLinkList{}
	single.ReverseList(head)
	single.PrintList(head)
}

// 递归玩法
func PrintReverse2(head Head) {
	if head == nil {
		return
	}
	PrintReverse2(head.Next)
	fmt.Printf("%v,", head.Data)
}

// 找出倒数第k的节点
// 后启动指针
func FindLastK(head Head, k int) *Node {
	if head == nil {
		return nil
	}
	cur, kcur := head, head
	n := 1
	for cur != nil {
		cur = cur.Next
		if n > k {
			kcur = kcur.Next
		}
		n++
	}
	// 压根没移动
	if kcur == head {
		return nil
	}
	return kcur
}

// 判断是否有环,并且找到入口
func FindRoundEntry(head Head) (*Node, bool) {
	if head == nil {
		return nil, false
	}
	cur := head
	m := make(map[*Node]struct{}) // 利用hashMap
	for cur != nil {
		if _, ok := m[cur]; ok {
			return cur, ok
		}
		m[cur] = struct{}{}
		cur = cur.Next
	}
	return nil, false
}

// 数学公式找环 2*(len+arc) = len+ n*per+arc ==>   len+arc=n*per ==> len=n*per-arc
// len:头部到环口 per:环周长 arc:快指针环偏移环口
// 根据简化的等式知: 从arc开始移动,那len=n*per,所以就能在环口相遇
// 快慢指针相遇之后,从head/slow一起继续走，就会在环口相遇
func FindRoundEntry2(head Head) (*Node, bool) {
	if head == nil {
		return nil, false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 无环
	if fast == nil || fast.Next == nil {
		return nil, false
	}
	// 有环,从slow/head一起走就能找到入口
	// 数学公式保证
	for slow != head {
		slow = slow.Next
		head = head.Next
	}
	return slow, true
}

// 判断汇聚单链表
// Deprecated: golang的map遍历是随机的,所以环内节点,都会在map中,方法不行！
// 采用2的方法吧!
func FindCrossEntry(head1, head2 Head) (*Node, bool) {
	if head1 == nil || head2 == nil {
		return nil, false
	}
	m1, m2 := make(map[*Node]struct{}), make(map[*Node]struct{})
	cur1, cur2 := head1, head2
	for cur1 != nil {
		m1[cur1] = struct{}{}
		cur1 = cur1.Next
	}
	for cur2 != nil {
		m2[cur2] = struct{}{}
		cur2 = cur2.Next
	}
	for k := range m1 {
		if _, ok := m2[k]; ok {
			return k, true
		}
	}
	return nil, false
}

// 也是数学,长的比短先走间隔的步数
// 然后同时走就能相遇了
func FindCrossEntry2(head1, head2 Head) (*Node, bool) {
	if head1 == nil || head2 == nil {
		return nil, false
	}
	single := SingleLinkList{}
	n1 := single.LengthList(head1)
	n2 := single.LengthList(head2)
	if n1 <= 0 || n2 <= 0 {
		return nil, false
	}
	diff := n1 - n2
	f1 := true
	if n2 > n1 {
		f1 = false
		diff = n2 - n1
	}
	cur1, cur2 := head1, head2
	// 走掉diff步
	if f1 {
		for cur1 != nil && diff > 0 {
			cur1 = cur1.Next
			diff--
		}
	} else {
		for cur2 != nil && diff > 0 {
			cur2 = cur2.Next
			diff--
		}
	}
	// 同时走
	for cur1 != nil && cur2 != nil && cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil || cur2 == nil {
		return nil, false
	}
	return cur1, true
}

// [a1,a2,a3,a4...an] -> [a1,an,a2,an-1,a3,an-2...]
// 空间O(n)算法,用arr
func HeadWithTail(head Head) {
	if head == nil {
		return
	}
	single := SingleLinkList{}
	var arr []*Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil // 需要把Next搞掉,不然会影响PushBack
		if cur.Data != nil {
			arr = append(arr, cur)
		}
		cur = next
	}
	n := len(arr)
	l, r := 0, len(arr)-1
	arr2 := make([]*Node, n)
	k := 0
	for l < r {
		if k+1 < n {
			arr2[k], arr2[k+1] = arr[l], arr[r]
		} else {
			arr2[k] = arr[l]
		}
		l++
		r--
		k += 2
	}
	head2 := single.InitList()
	for k := range arr2 {
		_ = single.PushBack(head2, arr2[k])
	}
	head.Next = head2.Next // 把head重新指向新的
	//head = head2 // 注意这里是有问题的！,return head2没问题,但是head是值传递,改变不了外面的，但是原来的head.Next已经nil了。(上面的步骤)
}

// 中间后面的转置
// 两个链表合并一下
// 虽然是"两个链表",但是空间复杂度是O(1) ps:区别一下数组
func HeadWithTail2(head Head) {
	n := SingleLinkList{}.LengthList(head)
	if n <= 0 {
		return
	}
	mid := n/2 + 1
	cur := head
	move := 1
	for cur.Next != nil {
		if move >= mid {
			break
		}
		cur = cur.Next
		move++
	}
	// 后半段的链表
	head2 := SingleLinkList{}.InitList(cur.Next)
	// 前后分离
	cur.Next = nil
	// 后半段转置
	SingleLinkList{}.ReverseList(head2)
	final := SingleLinkList{}.InitList()
	cur1, cur2 := head.Next, head2.Next // 需要无头链表
	left := true
	// 归并排序思想啦,只是不用比较,平均一人一次,注意需要无头链表
	for cur1 != nil && cur2 != nil {
		if left {
			next := cur1.Next
			cur1.Next = nil // 需要把Next干掉,不然PushBack有问题
			_ = SingleLinkList{}.PushBack(final, cur1)
			cur1 = next
			left = false
		} else {
			next := cur2.Next
			cur2.Next = nil
			_ = SingleLinkList{}.PushBack(final, cur2)
			cur2 = next
			left = true
		}
	}
	for cur1 != nil {
		_ = SingleLinkList{}.PushBack(final, cur1)
		cur1 = cur1.Next
	}
	for cur2 != nil {
		_ = SingleLinkList{}.PushBack(final, cur2)
		cur2 = cur2.Next
	}
	head2.Next = final.Next
}
