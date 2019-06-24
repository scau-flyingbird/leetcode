package minStack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStackExample1(t *testing.T) {
	minStack := &MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, minStack.GetMin(), -3, "should be equal")
	minStack.Pop()
	assert.Equal(t, minStack.Top(), 0, "should be equal")
	assert.Equal(t, minStack.GetMin(), -2, "should be equal")
}
