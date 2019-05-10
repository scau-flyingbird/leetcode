#### 206. Reverse Linked List

Reverse a singly linked list.

Example:


```
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

##### 算法思路：
1. 用ListNode pre,cur,nex分别前一个，当前，下一个节点；
2. 遍历一个一个翻转，翻转核心：

```
nex = cur.next;
cur.next = pre;
pre = cur;
cur = nex;
```
3. 如此类推过去，返回pre即为翻转结果。
###### 复杂度：O(n).

##### 代码：

```
package reverseLinkedList;

import model.ListNode;

public class ReverseLinkedList {
    public ListNode reverseList(ListNode head) {
        ListNode cur = head;
        ListNode pre = null;
        ListNode nex = null;
        while(cur != null){
            nex = cur.next;
            cur.next = pre;
            pre = cur;
            cur = nex;
        }
        return pre;
    }
}

```


