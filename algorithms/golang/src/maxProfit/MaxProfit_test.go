package maxProfit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfitExample1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 7, maxProfit(prices), "should be equal")
}
