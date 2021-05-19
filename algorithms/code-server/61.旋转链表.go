/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    n := 0
    for node := head; node != nil; n++ {
        node = node.Next
    }
    k = k % n
    if k == 0 {
        return head
    }
    
    A, tail := head, head
    for i := 0; i < k; i++ {
        tail = tail.Next
    }

    for tail.Next != nil {
        tail = tail.Next
        A = A.Next
    }

    ret := A.Next
    A.Next = nil
    tail.Next = head
    return ret
}
// @lc code=end

