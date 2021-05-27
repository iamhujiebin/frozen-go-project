package 桥接Bridge

import "testing"

func TestNewShape(t *testing.T) {
	shape := NewShape(&Rectangle{}, &Red{})
	shape.Draw()
	shape = NewShape(&Circle{}, &Green{})
	shape.Draw()
}
