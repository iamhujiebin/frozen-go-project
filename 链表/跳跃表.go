package 链表

import (
	"math/rand"
	"time"
)

// 跳跃表
// 实际上就是多层"有序链表"
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
type ISkipList interface {
	Insert(index uint64, data interface{})
	Delete(index uint64)
	Search(index uint64) *SkipListNode // key用于排序
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
		p:     0.25,
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

func (s *SkipList) Insert(index uint64, data interface{}) {
	panic("implement me")
}

func (s *SkipList) Delete(index uint64) {
	panic("implement me")
}

func (s *SkipList) Search(index uint64) *SkipListNode {
	panic("implement me")
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
	// 已经到了level=0的底层了,cur要不就是index节点,要不就是小于index的最小节点
	if cur.nextNodes[0] != s.tail {
		cur = cur.nextNodes[0]
	}
	return cur, preNodes
}

// 根据概率随机层高
func (s *SkipList) randLevel() int {
	rand.Seed(time.Now().UnixNano())
	level := 1
	for level < s.Level && rand.Float64() < s.p {
		level++
	}
	return level
}
