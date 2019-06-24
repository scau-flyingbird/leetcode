##### 217. Contains Duplicate

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

```
Example 1:
Input: [1,2,3,1]
Output: true

Example 2:
Input: [1,2,3,4]
Output: false

Example 3:
Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
```

##### 算法思路：
1. 用map[int]bool遍历存储nums数组,如果map中已存在当前num，则返回true；遍历完后没相同返回false。
##### 代码：

```
func containsDuplicate(nums []int) bool {
	intMap := map[int] bool{}
	for _, num := range  nums{
		_, ok := intMap[num]
		if ok {
			return true
		} else{
			intMap[num] = true
		}
	}
	return false
}
```

