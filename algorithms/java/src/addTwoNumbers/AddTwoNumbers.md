##### 2. Add Two Numbers

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```
Example 1:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

##### 算法思路：
1. 模拟列竖式加法运算，核心计算如下：

   ```
   int sum = (num1 + num2 + up);
   int val = sum%10;
   int up = sum/10;
   
   num1为l1.val，num2为l2.val，up为上次加得进位。
   ```

   

2. 链表连接代码：

   ```
   ListNode tmp = new ListNode(val);
   pre.next = tmp;
   pre = pre.next;
   ```

   

3. 遍历小技巧，while(l1 != null || l2 != null)，如果l1或l2为空，只需要把空的对应值当做0就可以，加法+0不影响结果。

4. 最后记得把进位也加上。

   ```
   if(up > 0){
   	pre.next = new ListNode(up);
   	pre = pre.next;
   }
   ```

   
##### 代码：

```
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


```
