package 状态State

import "fmt"

// 状态模式
// 把状态抽象成对象,对外方法不变,内部实现可修改
// 例子:机器的开关
type State interface {
	On()
	Off()
}

var (
	ONSTATE  = &ON{}
	OFFSTATE = &OFF{}
)

type ON struct {
}

func (O *ON) On() {
	fmt.Println("已经开着呢")
}

func (O *ON) Off() {
	fmt.Println("从On到Off")
}

type OFF struct {
}

func (O *OFF) On() {
	fmt.Println("从Off到ON")
}

func (O *OFF) Off() {
	fmt.Println("已经关着呢")
}

type Machine struct {
	state State
}

func (m *Machine) SetState(state State) {
	m.state = state
}

func (m *Machine) ON() {
	m.state.On()
	m.SetState(ONSTATE)
}

func (m *Machine) OFF() {
	m.state.Off()
	m.SetState(OFFSTATE)
}

func NewMachine() *Machine {
	return &Machine{state: OFFSTATE}
}
