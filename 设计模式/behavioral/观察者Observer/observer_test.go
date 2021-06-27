package 观察者Observer

import (
	"fmt"
	"testing"
)

func TestConcreteSubject_Notify(t *testing.T) {
	subject := NewConcreteSubject()
	o1 := NewConcreteObserver("observer1")
	o2 := NewConcreteObserver("observer2")
	o3 := NewConcreteObserver("observer3")
	subject.AddListener(o1)
	subject.AddListener(o2)
	subject.AddListener(o3)
	for x := range Fib(10) {
		subject.Notify(Event{
			state: fmt.Sprintf("state:%d", x),
		})
		subject.RemoveListener(o3)
	}
}
