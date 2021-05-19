/*
 * @lc app=leetcode.cn id=1503 lang=golang
 *
 * [1503] 所有蚂蚁掉下来前的最后一刻
 */

// @lc code=start
func max(a, b int) int{
    if a > b {
        return a
    }else {
        return b
    }
}

func getLastMoment(n int, left []int, right []int) int {
    ret := 0

    for _, num := range left {
        ret = max(ret, num)
    }

    for _, num := range right {
        ret = max(ret, n - num)
    }

    return ret
}
// @lc code=end

