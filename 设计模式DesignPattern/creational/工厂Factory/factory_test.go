package 工厂Factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPayChannel(t *testing.T) {
	google, err := NewPayChannel("google")
	assert.Equal(t, nil, err)
	google.CreateOrder()
	facebook, err := NewPayChannel("facebook")
	assert.Equal(t, nil, err)
	facebook.CreateOrder()
	wrong, err := NewPayChannel("wrong")
	assert.Equal(t, wrong, nil)
	assert.Equal(t, true, err != nil)
}
