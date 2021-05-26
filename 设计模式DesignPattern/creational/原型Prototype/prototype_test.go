package 原型Prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// golang判断两个对象是否相等,判断类型+存储空间+值
func TestConcretePrototype_Clone(t *testing.T) {
	src := NewConcretePrototype()
	dest := src.Clone()
	assert.Equal(t, src.Name(), dest.Name())
	assert.Equal(t, false, src == dest) // src dest的地址不同,内存空间不同
}
