/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
    cur, n := 0, len(s)
    str := []byte(s)

    for ; cur < n ; {
        left := cur
        for ;cur < n && s[cur] != ' '; cur++ {

        }
        right := cur - 1
        for left < right {
            str[left], str[right] = str[right], str[left]
            left++
            right--
        }
        cur++
    }
    return string(str)
}
// @lc code=end

