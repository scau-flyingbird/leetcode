package intersectionofTwoLinkedLists;

import model.ListNode;

import java.util.ArrayList;
import java.util.List;

public class IntersectionofTwoLinkedLists {

    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        List<Integer> listA = new ArrayList<Integer>();
        List<Integer> listB = new ArrayList<Integer>();
        while(headA != null){
            listA.add(headA.val);
            headA = headA.next;
        }
        while(headB != null){
            listB.add(headB.val);
            headB = headB.next;
        }
        int i = listA.size() - 1,j = listB.size() - 1;
        for(; i >=0 && j>=0; i--,j--){
            if(listA.get(i) != listB.get(j)){
                break;
            }
        }
        if(listA.size() - i == 0){
            return null;
        }
        else{
            ListNode head = new ListNode(i+1);
            ListNode node = head;
            for(int k = i + 2; k < listA.size(); k++){
                node.next = new ListNode(listA.get(k));
                node = node.next;
            }
            return head;
        }
    }
}
