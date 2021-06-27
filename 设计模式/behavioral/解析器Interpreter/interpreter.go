package 解析器Interpreter

import "strings"

// 解析器模型
// 语法数解析,把有规律的操作抽象出来
// 例子: 写一个实现加法的解析器
// 这个例子只是解决两个数的相加,不支持三个以上.以为是大佬,结果很蠢
type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Integer struct {
	integer int
}

func (n *Integer) Interpret(variables map[string]Expression) int {
	return n.integer
}

type Plus struct {
	leftExpression  Expression
	rightExpression Expression
}

func (p *Plus) Interpret(variables map[string]Expression) int {
	return p.leftExpression.Interpret(variables) + p.rightExpression.Interpret(variables)
}

type Variable struct {
	Name string
}

func (v *Variable) Interpret(variables map[string]Expression) int {
	value, found := variables[v.Name]
	if !found {
		return 0
	}
	return value.Interpret(variables)
}

type Evaluator struct {
	syntaxTree Expression
}

func (p *Evaluator) Interpret(variables map[string]Expression) int {
	return p.syntaxTree.Interpret(variables)
}

// 搞一个Stack
// 暂时不搞线程安全
type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{
		value: value,
		next:  s.top,
	}
	s.size++
}

func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		switch token {
		case "+":
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		default:
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree: syntaxTree}
}
