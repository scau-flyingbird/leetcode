/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}

	p, q, ret := nums[0], max(nums[0], nums[1]), 0
	for i := 2; i < n; i++ {
		ret = max(p + nums[i], q)
		p = q
		q = ret
	}
	return ret
}

// @lc code=end

