/*
 * @lc app=leetcode.cn id=1033 lang=golang
 *
 * [1033] 移动石子直到连续
 */

// @lc code=start
func numMovesStones(a int, b int, c int) []int {
    if a > b {
        a, b = b, a
    }
    if a > c {
        a, c = c, a
    }
    if b > c {
        b, c = c, b
    }

    maxCnt := b - (a + 1) + c - (b + 1)

    minCnt := 0
    if b == a + 1 && c == b + 1 {
        minCnt = 0
    }else if b - a <= 2 || c - b <= 2 {
        minCnt = 1
    }else {
        minCnt = 2
    }

    return []int{minCnt, maxCnt}
}
// @lc code=end

