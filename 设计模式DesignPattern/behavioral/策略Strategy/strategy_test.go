package 策略Strategy

import "testing"

func TestNewContext(t *testing.T) {
	c := NewContext(&PlanAStrategy{})
	c.Execute()
	c.SetStrategy(&PlanBStrategy{})
	c.Execute()
}
