package 状态State

import "testing"

func TestNewMachine(t *testing.T) {
	m := NewMachine()
	m.ON()
	m.ON()
	m.OFF()
	m.OFF()
}
