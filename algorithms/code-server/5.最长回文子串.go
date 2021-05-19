/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	idx, maxLen := 0, 1
	n := len(s)

	dp := make([][]bool, n)
	for i := 0 ;i < n ;i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for len := 2; len <= n; len++ {
		for i := 0 ; i + len <= n; i++ {
			j := i + len - 1
			if s[i] != s[j] {
				continue
			}
			if len > 2 && !dp[i+1][j-1] {
				continue
			}
			dp[i][j] = true
			if len > maxLen {
				idx = i
				maxLen = len
			}
		}
	}
	return string(s[idx: (maxLen + idx)])
}
// @lc code=end

