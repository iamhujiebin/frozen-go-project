package 构造者Buider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDirector(t *testing.T) {
	builder := NewConcreteBuilder()
	product := NewDirector(builder).Construct()
	assert.Equal(t, product.Part1 == "build part1", true)
	assert.Equal(t, product.Part2 == "build part2", true)

	assert.Equal(t, builder.GetResult(), product) // 所有变量的值都相等
	assert.Equal(t, builder.GetResult() == product, false)
}
