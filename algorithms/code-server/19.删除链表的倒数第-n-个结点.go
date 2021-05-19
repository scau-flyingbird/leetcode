/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    var A, B, K *ListNode
    A, B, K = nil, head, head

    for i := 0 ;i < n ; i++ {
        K = K.Next
    }

    for K != nil {
        K = K.Next
        A = B
        B = B.Next
    }
    
    if A != nil {
        A.Next = B.Next
        return head
    }else {
        return B.Next
    }
}
// @lc code=end

