package 享元FlyWeight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFlyWeightFactory(t *testing.T) {
	factory := NewFlyWeightFactory()
	one := factory.GetFlyWeight("one")
	two := factory.GetFlyWeight("one")

	assert.Equal(t, one == two, true) // 指针相同,同一个对象
}
