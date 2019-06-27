package trappingRainWater

/**
双指针算法：由DP算法优化空间可得
leftMax[]int,rightMax[]int 每个值用过一次就没用了，其实就只用leftMax int,rightMax int保存左右两边柱子最大值即可。
而leftMax = Math.max(leftMax, height[left-1]),rightMax = Math.max(rightMax, height[right+1])得到的，
由此可知，
①if height[left-1] < height[right+1] ,只对于当前left柱子来说，
可装的水就肯定是 if leftMax > height[left] ,可装水tmp =  leftMax - height[left], sum+=tmp，这时left柱子已经计算了，则left++；
②if height[left-1] >= height[right+1]，只对于当前right柱子来说，
可装的水就肯定是 if rightMax > height[right]，可装水tmp = rightMax - height[right]，sum+=tmp，这时right柱子已经计算了，则right--；
一直循环到left<=right为false
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
动态规划：leftMax[]int,rightMax[]int
存储leftMax[i]从0-i最大的柱子（左）/rightMax[i]从i-[len(height)-1]最大的柱子（右）
当 min := getMin(leftMax[i-1], rightMax[i+1])/左右两边最大柱子的较少值柱子高度
	min > height[cur] 时，能装水，数量为min-height[cur]
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


