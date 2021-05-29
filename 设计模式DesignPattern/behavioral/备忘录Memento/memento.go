package 备忘录Memento

// 备忘录模式
// 记录某个对象的某个时候的状态,可用于恢复. 游戏/数据库redo/undo之类
// 三个对象: Memento:备忘录 Originator:发起人 CareTaker: 备忘录管理者
type Memento struct {
	State string
}

type Originator struct {
	state string
}

func (p *Originator) Init() {
	p.state = "init"
}

func (p *Originator) Grow() {
	p.state = "grow"
}

func (p *Originator) Old() {
	p.state = "old"
}

func (p *Originator) State() string {
	return p.state
}

func (p *Originator) CreateMemento() (memento *Memento) {
	return &Memento{State: p.state}
}

func (p *Originator) RestoreState(memento *Memento) {
	p.state = memento.State
}

type CareTaker struct {
	Memento *Memento
}

func (c *CareTaker) SetMemento(memento *Memento) {
	c.Memento = memento
}

func (c *CareTaker) GetMemento() *Memento {
	return c.Memento
}

func NewOriginator() *Originator {
	return &Originator{state: "create"}
}
