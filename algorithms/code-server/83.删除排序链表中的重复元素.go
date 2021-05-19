/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := head

    for cur != nil {
        if cur.Next != nil && cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
        }else {
            cur = cur.Next
        }
    }
    
    return head
}
// @lc code=end

