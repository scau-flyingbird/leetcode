package trappingRainWater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrapExample1(t *testing.T) {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	assert.Equal(t, 6, trap(height), "should be equal")
}