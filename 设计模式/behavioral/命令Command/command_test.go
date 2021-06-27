package 命令Command

import "testing"

func TestNewInvoker(t *testing.T) {
	tv := &TVReceiver{Name: "xiaomi tv"}
	ch1Cmd := NewChannel1Command(tv)
	invoker := NewInvoker(ch1Cmd)
	invoker.Call()

	invoker.SetCmd(NewChannel2Command(tv))
	invoker.Call()
}
