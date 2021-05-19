/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum, max, n := 0, 0, len(nums)
	for _,num := range nums{
		sum += num
		if num > max {
			max = num
		}
	}

	target := sum/2
	if sum % 2 != 0 || max > target {
		return false
	}

	dp := make([]bool, target + 1)
	dp[0] = true
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := target; j >= 0; j-- {
			if j >= v {
				dp[j] = dp[j] || dp[j-v]
			}
		}
	}

	return dp[target]
}
// @lc code=end

