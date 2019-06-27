
##### 42. Trapping Rain Water
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

![https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png]()


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

**Example:**

```  
Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```
##### 算法思路：  
1.动态规划：

**①数据结构：**leftMax[]int,rightMax[]int 存储leftMax[i]从0-i最大的柱子（左）/rightMax[i]从i-[len(height)-1]最大的柱子（右）；
**②动态规划数组：**leftMax,rightMax为动态规划数组，leftMax[i]=Math.max(height[i],leftMax[i-1]);
**③核心算法：**

当前节点i，左右两边最大柱子较少的值min,则可装水tmp=min-height[i]。
min := getMin(leftMax[i-1], rightMax[i+1])/左右两边最大柱子的较少值柱子高度。
min > height[i] 时，能装水，数量为min-height[i]。

2.双指针优化动态规划：

**①算法思路：**由DP算法优化空间可得。
**②优化点：**leftMax[]int,rightMax[]int 每个值用过一次就没用了，只用**leftMax int,rightMax int**保存左右两边柱子最大值即可。
**③核心思路：**
leftMax = Math.max(leftMax, height[left-1]),
rightMax = Math.max(rightMax, height[right+1])
可推理的（基础：左右两边最大柱子的较少值柱子高度）：
2.1 if height[left-1] < height[right+1] ,只对于当前left柱子来说，左右两边最大的柱子较少值肯定是左边leftMax ，可装的水就肯定是 if leftMax > height[left] ,可装水tmp =  leftMax - height[left], sum+=tmp，这时left柱子已经计算了，则left++；
2.2 if height[left-1] >= height[right+1]，只对于当前right柱子来说，左右两边最大的柱子较少值肯定是右边rightMax ，可装的水就肯定是 if rightMax > height[right]，可装水tmp = rightMax - height[right]，sum+=tmp，这时right柱子已经计算了，则right--；
**④一直循环到left<=right为false。**



##### 代码：  

```  
package trappingRainWater

/**
双指针
 */
func trap(height []int) int {
	left := 1
	right := len(height) - 2
	maxLeft := 0
	maxRight := 0
	sum := 0
	for ; left <= right ; {
		if height[left - 1] < height [right + 1] {
			maxLeft = getMax(maxLeft, height[left - 1])
			min := maxLeft
			if min > height[left] {
				sum += min - height[left]
			}
			left++
		}else {
			maxRight = getMax(maxRight, height[right+1])
			min := maxRight
			if min > height[right] {
				sum += min - height[right]
			}
			right--
		}
	}
	return sum
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}


/**
动态规划
 */
func trap1(height []int) int {
	leftMax := getLeftMax(height)
	rightMax := getRightMax(height)
	sum := 0
	for i := 1 ;i < len(height) - 1; i++ {
		min := getMin(leftMax[i-1], rightMax[i+1])
		if height[i] < min {
			sum += min - height[i]
		}
	}
	return sum
}

func getMin(left int, right int) int {
	if left < right {
		return left
	}else {
		return right
	}
}

func getLeftMax(height []int) []int{
	leftMax := make([]int, len(height))
	if len(height) == 0 {
		return leftMax
	}
	leftMax[0]=height[0]
	for i:=1; i < len(height); i++ {
		if leftMax[i-1] < height[i] {
			leftMax[i] = height[i]
		}else {
			leftMax[i] = leftMax[i-1]
		}
	}
	return leftMax
}

func getRightMax(height []int) []int{
	rightMax := make([]int, len(height))
	if len(height) == 0 {
		return rightMax
	}
	rightMax[len(height)-1]=height[len(height)-1]
	for i:=len(height)-1; i >= 0; i-- {
		if rightMax[i+1] < height[i] {
			rightMax[i] = height[i]
		}else {
			rightMax[i] = rightMax[i+1]
		}
	}
	return rightMax
}
```
