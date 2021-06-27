package 抽象工厂AbstractFactory

import "testing"

func TestNewSimpleLunch(t *testing.T) {
	NewSimpleLunch().CreateFood().Cook()
	NewSimpleLunch().CreateVegetable().Cook()
}
