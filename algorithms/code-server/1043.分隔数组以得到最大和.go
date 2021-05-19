/*
 * @lc app=leetcode.cn id=1043 lang=golang
 *
 * [1043] 分隔数组以得到最大和
 */

// @lc code=start
func max(a, b int) int{
	if a > b {
		return a
	}else {
		return b
	}
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
    f := make([]int, n + 1)
    for i := 1; i <= n; i++ {
		//[1, i]
		mmax := math.MinInt32
		for j := i; j > max(i - k, 0); j-- {
			//[1,i],[1]
			mmax = max(mmax, arr[j-1])
            f[i] = max(f[i], f[j - 1] + mmax * (i - j + 1))
        }
    }
    return f[n]
}
// @lc code=end

