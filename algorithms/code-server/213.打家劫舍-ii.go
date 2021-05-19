/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func max(a, b int) int{
	if a > b {
		return a
	}else {
		return b
	}
}

func robV1(nums []int) int{
	n := len(nums)
	if n <= 0 {
		return 0
	}
	p, q, ret := 0, 0, 0
	for i := 0; i < n; i++ {
		ret = max(p + nums[i], q)
		p = q
		q = ret
	}
	return ret
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}else if len(nums) == 1 {
		return nums[0]
	}
	return max(robV1(nums[1:]), robV1(nums[:len(nums)-1])) 
}
// @lc code=end

