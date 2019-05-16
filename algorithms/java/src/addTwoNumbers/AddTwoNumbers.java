package addTwoNumbers;

import model.ListNode;

public class AddTwoNumbers {

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode l = null;
        ListNode head = null;
        int n = 0;
        int up = 0;
        while(l1 != null || l2 != null){
            int sum = 0;
            if(l1 == null){
                sum = up + l1.val;
                up = sum/10;
                sum = sum%10;
            }
            else if(l2 == null){
                sum = up + l2.val;
                up = sum/10;
                sum = sum%10;
            }else{
                sum = up + l1.val + l2.val;
                up = sum/10;
                sum = sum%10;
            }
            if(l != null){
                l.next = new ListNode(sum);
            }
            else {
                l = new ListNode(sum);
                head = l;
            }
        }
        if(up > 0){

        }
        return l;
    }
}
