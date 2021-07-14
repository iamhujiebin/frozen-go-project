package 链表

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 跳跃表
// 实际上就是多层"有序链表",理解也是当成"单链表"来理解就行
// 排序根据index字段
// 层高level,由节点数决定
//	log(1/PROBABILITY)(N),PROBABILITY 上升一层概率,比如PROBABILITY=0.25,那一个节点有两层的概率是1/4,有三层的概率1/16,有四层的概率1/64....
//  log(1/0.25)(N) = log(4)N ==> log(4)10000000 ≈ 12
// 关键要素
// 1. 层高Level
// 2. 头节点Head、尾节点Tail
// 3. P(PROBABILITY):概率,一般是0.5 / 0.25
// 4. NextNodes数组
// 	 4.1 level多少,数组元素就有多少个
//   4.2 类比一下链表,跳跃表也有头指针head(没有具体的index/data),但是有Next(数组)
// 5. 操作
//	5.1 插入
//  5.2 删除
//  5.3 查找! 重点,插入、删除都依赖查找,查找利用概率论,可以跟平衡二叉树抗衡 O(logN)
// 图理解,请看Search方法
type ISkipList interface {
	Insert(index uint64, data interface{})
	Delete(index uint64)
	Search(index uint64) *SkipListNode // key用于排序
	Range() []*SkipListNode            // 遍历
}

// 跳跃表节点
type SkipListNode struct {
	Index     uint64
	Data      interface{}
	nextNodes []*SkipListNode
}

type SkipList struct {
	level int
	head  *SkipListNode
	tail  *SkipListNode // 都是空
	p     float64
}

// 创建新的跳跃表
func NewSkipList(level int) *SkipList {
	head := newSkipListNode(0, nil, level)
	var tail *SkipListNode
	for i := range head.nextNodes {
		head.nextNodes[i] = tail
	}
	return &SkipList{
		head:  head,
		tail:  tail,
		level: level,
		p:     0.5,
	}
}

// 创建新的跳跃表节点
func newSkipListNode(index uint64, data interface{}, level int) *SkipListNode {
	return &SkipListNode{
		Index:     index,
		Data:      data,
		nextNodes: make([]*SkipListNode, level),
	}
}

// 看着图来理解.level=4
/*
Head.nextNode[1]->1⃣️----->3⃣️----------------------------->Tail
Head.nextNode[1]->1⃣️----->3⃣️--------------------->9⃣️->🔟->Tail
Head.nextNode[1]->1⃣️----->3⃣️->4⃣️->5⃣️->6⃣️--------->9⃣️->🔟->Tail
Head.nextNode[0]->1⃣️->2⃣️->3⃣️->4⃣️->5⃣️->6⃣️->7⃣️->8⃣️->9⃣️->🔟->Tail
*/
// 注意:1⃣️~🔟,都是"同一个"节点,指针地址是一样的。都有nextNode[4],如果哪层的后面节点为空,就是nextNode[l] = Tail
// 举个🌰:查找index=10
func (s *SkipList) Search(index uint64) *SkipListNode {
	cur := s.head
	// 从上层往下找
	for l := s.level - 1; l >= 0; l-- {
		for cur.nextNodes[l] != s.tail && cur.nextNodes[l].Index < index {
			cur = cur.nextNodes[l]
		}
		if cur.nextNodes[l] != s.tail && cur.nextNodes[l].Index == index {
			return cur.nextNodes[l]
		}
	}
	// 找到的level=0,都没有index相等的
	return nil
}

// 返回搜索节点(可能是搜索节点的前驱节点)以及各层的前驱节点
// 用于插入/删除新的节点
// 前驱节点,是数组,是每层链表对应节点的前驱,所有层都要,因为插入
func (s *SkipList) searchWithPre(index uint64) (*SkipListNode, []*SkipListNode) {
	preNodes := make([]*SkipListNode, s.level)
	cur := s.head
	// 从上层往下找
	for l := s.level - 1; l >= 0; l-- {
		for cur.nextNodes[l] != s.tail && cur.nextNodes[l].Index < index {
			cur = cur.nextNodes[l]
		}
		// 跳出循环:要不到层的末尾,要不找到了一个index大的,cur就是前一个节点
		preNodes[l] = cur
	}
	// 已经到了level=0的底层了,cur就是小于index的最大节点
	if cur.nextNodes[0] != s.tail {
		cur = cur.nextNodes[0]
	}
	return cur, preNodes
}

// Deprecated: 废弃
func (s *SkipList) SearchNotGood(index uint64) *SkipListNode {
	// 顶层的next开始,"不能"直接指向next,因为上层找不到,不一定下层没有,cur不能这么快后移
	// 后移前,需要先判断cur.nextNodes[l]
	cur := s.head.nextNodes[s.level-1]
	// 从上层往下找
	for l := s.level - 1; l >= 0; l-- {
		for cur != s.tail {
			if cur.Index == index {
				return cur
			}
			cur = cur.nextNodes[l]
		}
	}
	return nil
}

func (s *SkipList) Insert(index uint64, data interface{}) {
	// 找前驱节点
	cur, preNodes := s.searchWithPre(index)
	if cur != s.head && cur.Index == index {
		cur.Data = data
		return
	}
	newNode := newSkipListNode(index, data, s.randLevel())
	level := len(newNode.nextNodes) // 新节点的层高
	for l := level - 1; l >= 0; l-- {
		// 链表插入
		newNode.nextNodes[l] = preNodes[l].nextNodes[l]
		preNodes[l].nextNodes[l] = newNode
	}
}

func (s *SkipList) Delete(index uint64) {
	// 找前驱节点
	cur, preNodes := s.searchWithPre(index)
	if cur != s.head && cur.Index != index {
		// 没有对应的节点
		return
	}
	level := len(cur.nextNodes)
	for l := level - 1; l >= 0; l-- {
		// 链表删除
		preNodes[l].nextNodes[l] = cur.nextNodes[l]
	}
}

func (s *SkipList) Range() []*SkipListNode {
	if s.head == nil {
		return nil
	}
	var nodes []*SkipListNode
	// 只遍历最后一层就行了
	l := 0
	cur := s.head.nextNodes[l]
	for cur != s.tail {
		nodes = append(nodes, cur)
		cur = cur.nextNodes[l]
	}
	return nodes
}

// 根据概率随机层高
func (s *SkipList) randLevel() int {
	level := 1
	for level < s.level && rand.Float64() < s.p {
		level++
	}
	return level
}
