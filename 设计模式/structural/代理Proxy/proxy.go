package 代理Proxy

import "fmt"

// 顾名思义就好,中间商做事情
// Agent要做的事情定义，跟实际去做的Task定义一样
type ITask interface {
	RentHouse(addr string, price float64)
}

func NewConcreteTask() *Task {
	return &Task{}
}

type Task struct {
}

func (t *Task) RentHouse(addr string, price float64) {
	fmt.Printf("%s 的房子 ￥%.2f 钱,中介费 ￥%.2f\n", addr, price, price*0.02)
}

type AgentTask struct {
	task ITask
}

func (p *AgentTask) RentHouse(addr string, price float64) {
	p.task.RentHouse(addr, price)
}

func NewAgentTask(task ITask) *AgentTask {
	return &AgentTask{task: task}
}
