package 适配器Adapter

import "testing"

func TestAdaptee_SpecialExecute(t *testing.T) {
	NewAdapter(&Adaptee{}).Execute()
}
