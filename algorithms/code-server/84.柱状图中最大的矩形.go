/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0)
	width := make([]int, n)
	for i := 0 ; i < len(width); i++ {
		width[i] = 1
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			width[i] += i
		}else {
			width[i] += i - stack[len(stack) - 1] - 1
		}

		stack = append(stack, i)
	}

	ret := 0
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			width[i] += n - i - 1
		}else {
			width[i] += stack[len(stack) - 1] - i - 1
		}

		ret = max(ret, heights[i] * width[i])
		stack = append(stack, i)
	}
	return ret
}
// @lc code=end

