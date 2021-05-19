/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func min(a ,b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ret := 0
	stack := make([]int, 0)
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack) - 1]]{
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack) - 1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ret += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return ret
}
// @lc code=end

