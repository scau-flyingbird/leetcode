/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
var ret []string
var n int
var s []byte

func dfs(x int) {
    if x == n {
        ret = append(ret, string(s))
        return
    }

    dfs(x + 1)
    if s[x] >= 'a' && s[x] <= 'z' {
        tmp := s[x]
        s[x] = byte(s[x] - 'a' + 'A')
        dfs(x + 1)
        s[x] = tmp
    }else if s[x] >= 'A' && s[x] <= 'Z' {
        tmp := s[x]
        s[x] = byte(s[x] - 'A' + 'a')
        dfs(x + 1)
        s[x] = tmp
    }
}

func letterCasePermutation(S string) []string {
    s = []byte(S)
    n = len(s)
    ret = make([]string, 0)
    dfs(0)
    return ret
}
// @lc code=end

