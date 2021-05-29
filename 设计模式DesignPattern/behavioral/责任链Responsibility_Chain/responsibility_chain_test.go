package 责任链Responsibility_Chain

import "testing"

func TestNewDayOffRequest(t *testing.T) {
	handler := NewDayOffRequest()
	handler.HandlerRequest(1)
	handler.HandlerRequest(3)
	handler.HandlerRequest(7)
	handler.HandlerRequest(11)
}
