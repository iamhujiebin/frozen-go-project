package 备忘录Memento

import "testing"

func TestNewOriginator(t *testing.T) {
	originator := NewOriginator()
	originator.Init()
	originator.Grow()
	memento := originator.CreateMemento()
	originator.Old()
	println(originator.State())
	originator.RestoreState(memento)
	println(originator.State())
}
