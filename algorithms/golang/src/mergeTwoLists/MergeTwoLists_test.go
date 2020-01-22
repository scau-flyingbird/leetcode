package mergeTwoLists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoListsExample1(t *testing.T) {
	//1->2->4, 1->3->4
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	//1->1->2->3->4->4
	l3 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	assert.Equal(t, l3, mergeTwoLists(l1, l2), "should be equal")
}
