/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	bitmap := 0
	for _, num := range nums {
		bitmap ^= num
	}

	mask := bitmap & (-bitmap)
	a, b := 0, 0
	for _, num := range nums {
		if num & mask == 0 {
			a ^= num
		}else {
			b ^= num
		}
	}
	return []int{a, b}
}
// @lc code=end

