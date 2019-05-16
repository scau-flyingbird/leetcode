Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

##### Example:

```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```

##### 算法思路：
递归Merge：
1.当l1为空，则返回l2；
2.当l2为空，则返回l1;
3.当l1和l2都不为空，则判断：
- l1.val<=l2.val，则l1.next=递归Merge(l1.next,l2),返回l1即可；
- l1.val>l2.val，则l2.next=递归Merge(l1,l2.next),返回l2即可;

##### 代码：

```
package mergeTwoSortedLists;

import model.ListNode;

public class MergeTwoSortedLists {

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if(l1 == null){
            return l2;
        }
        if(l2 == null){
            return l1;
        }
        if(l1.val <= l2.val){
            l1.next = mergeTwoLists(l1.next,l2);
            return l1;
        }
        else{
            l2.next = mergeTwoLists(l1, l2.next);
            return l2;
        }
    }
}

```
