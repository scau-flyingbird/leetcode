/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    for l := 0 ; ; l++ {
        i := 0
        for ; i < n; i++ {
            if l >= len(strs[i]) || i > 0 && strs[i][l] != strs[i-1][l] {
                break
            }
        }
        if i != n {
            return strs[0][:l]
        }
    }
    return ""
}
// @lc code=end

