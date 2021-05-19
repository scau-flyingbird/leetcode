/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	ret := 1
	for ret <= num {
		ret <<= 1
	}
	return (ret - 1)^num
}
// @lc code=end

