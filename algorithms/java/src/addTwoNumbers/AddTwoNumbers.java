package addTwoNumbers;

import model.ListNode;

public class AddTwoNumbers {

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = new ListNode(0);
        ListNode pre = head;
        int up = 0;
        while(l1 != null || l2 != null){
            int num1 = l1 == null? 0:l1.val;
            int num2 = l2 == null? 0:l2.val;
            int sum = (num1 + num2 + up);
            int val = sum%10;
            up  = sum/10;
            ListNode tmp = new ListNode(val);
            pre.next = tmp;
            pre = pre.next;
            l1 = l1 == null? null:l1.next;
            l2 = l2 == null? null:l2.next;
        }
        if(up > 0){
            pre.next = new ListNode(up);
            pre = pre.next;
        }
        return head.next;
    }
}
