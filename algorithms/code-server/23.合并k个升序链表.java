import java.util.Comparator;
import java.util.PriorityQueue;

/*
 * @lc app=leetcode.cn id=23 lang=java
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {

    private PriorityQueue<ListNode> queue = new PriorityQueue<ListNode>((x, y) -> x.val - y.val);

    public ListNode mergeKLists(ListNode[] lists) {
            for(int i = 0; i < lists.length; i++){
                if(lists[i] != null) {
                    queue.offer(lists[i]);
                }
            }

            ListNode head = null;
            ListNode tail = null;
            while(!queue.isEmpty()){
                ListNode node = queue.poll();
                if(head == null){
                    head = node;
                    tail = head;
                }else {
                    tail.next = node;
                    tail = tail.next;
                }
                if(tail.next != null){
                    queue.offer(tail.next);
                }
            }
            return head;
    }
}
// @lc code=end

