package 解析器Interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvaluator(t *testing.T) {
	expression := "x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["x"] = &Integer{1}
	variables["z"] = &Integer{2}
	result := sentence.Interpret(variables)
	assert.Equal(t, 3, result)
}
