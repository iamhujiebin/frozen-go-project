package 命令Command

import "fmt"

// 命令模式
// 官方定义:"将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化，对请求排队或记录请求日志，以及支持可撤销的操作"
// 把命令抽出来成为对象,是调用者和接受者的桥梁
// 三个元素,Invoker,Receiver,Command
// 具体例子:遥控器控制电视转台,整个遥控器就是Invoker,具体按钮就是Command的实现,电视就是Receiver
type Command interface {
	Execute()
}

func NewInvoker(cmd Command) *Invoker {
	return &Invoker{
		cmd: cmd,
	}
}

type Invoker struct {
	cmd Command
}

// 执行命令
func (p *Invoker) Call() {
	p.cmd.Execute()
}

func (p *Invoker) SetCmd(cmd Command) {
	p.cmd = cmd
}

type TVReceiver struct {
	Name string
}

func NewChannel1Command(receiver *TVReceiver) *TVChannel1Command {
	return &TVChannel1Command{receiver: receiver}
}

type TVChannel1Command struct {
	receiver *TVReceiver
}

func (T *TVChannel1Command) Execute() {
	fmt.Printf("%s to channel 1\n", T.receiver.Name)
}

func NewChannel2Command(receiver *TVReceiver) *TVChannel2Command {
	return &TVChannel2Command{receiver: receiver}
}

type TVChannel2Command struct {
	receiver *TVReceiver
}

func (T *TVChannel2Command) Execute() {
	fmt.Printf("%s to channel 2\n", T.receiver.Name)
}
