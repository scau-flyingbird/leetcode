/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
    n, ret := len(height), 0
    for l, r := 0, n - 1; l < r ; {
        ret = max(min(height[l], height[r]) * (r - l), ret)
        if height[l] <= height[r] {
            l++
        }else {
            r--
        }
    }
    return ret
}

func min(a, b int) int {
    if a < b {
        return a
    }else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}
// @lc code=end

