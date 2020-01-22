##### 21. Merge Two Sorted Lists

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

```
Example 1:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```

##### 算法思路：
利用递归算法，判断当前入参l1,l2的值大小，设较小值List为Lmin，较大值List为Lmax，则应赋值Lmin.Next=递归算法(Lmin.Next,Lmax)，返回Lmin：

a)如果l1为空，则返回l2；

b)如果l2为空，则返回l1;

c)如果l1.val小于等于l2.val，则返回l1，且返回l1前，赋值l1.Next=递归算法(l1.Next, l2)；

d)如果l1.val大于l2.val，则返回l2，且返回l2前，赋值l2.Next=递归算法(l1, l2.Next)；

##### 代码：

```
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
```

