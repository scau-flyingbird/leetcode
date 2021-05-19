/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	ret := 0
	for i := 0 ;i < 64; i++ {
		cnt := 0
		for _, num := range nums {
			cnt += num >> i & 1
		}
		ret |= (cnt % 3) << i
	}
	return ret
}
// @lc code=end

