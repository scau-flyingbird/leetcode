/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] 最长特殊序列 Ⅰ
 */

// @lc code=start
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}else {
		return int(math.Max(float64(len(a)), float64(len(b))))
	}
}
// @lc code=end

