package containsDuplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicateExample1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	assert.Equal(t, true, containsDuplicate(nums), "should be equal")
}

func TestContainsDuplicateExample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, false, containsDuplicate(nums), "should be equal")
}

func TestContainsDuplicateExample3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	assert.Equal(t, true, containsDuplicate(nums), "should be equal")
}
