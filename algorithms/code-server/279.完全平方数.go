/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func min(a, b int) int {
    if a < b {
        return a
    }else {
        return b
    }
}

func numSquares(n int) int {
    f := make([]int, n + 1)
    for i := 0 ;i < n + 1; i++ {
        if i == 0 {
            f[i] = 0
        }else {
            f[i] = math.MaxInt32
        }
    }

    for i := 1 ;i * i <= n; i++ {
        f[i*i] = 1
    }

    for i := 1 ; i <= n; i++ {
        for j := 1; j * j < i; j++ {
            f[i] = min(f[i - j * j] + 1, f[i])
        }
    }

    return f[n]
}
// @lc code=end

