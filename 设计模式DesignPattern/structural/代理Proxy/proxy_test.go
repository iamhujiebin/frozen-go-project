package 代理Proxy

import "testing"

func TestNewAgentTask(t *testing.T) {
	NewAgentTask(NewConcreteTask()).RentHouse("广州", 4000000)
}
