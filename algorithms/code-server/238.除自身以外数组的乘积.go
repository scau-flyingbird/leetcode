/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1 ;i < n; i++ {
		ans[i] = nums[i - 1] * ans[i - 1]
	}

	R := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R = R * nums[i]
	}

	return ans
}
// @lc code=end

