/*
 * @lc app=leetcode.cn id=1340 lang=golang
 *
 * [1340] 跳跃游戏 V
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}
func maxJumps(arr []int, d int) int {
	ret, n := 0, len(arr)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	var dfs func(arr []int, i int, d int)
	dfs = func(arr []int, i int, d int) {
		if dp[i] != -1 {
			return 
		}

		dp[i] = 1
		for j := i - 1; j >= 0 && j >= i - d && arr[j] < arr[i]; j-- {
			dfs(arr, j, d)
			dp[i] = max(dp[i], dp[j] + 1)
		}

		for j := i + 1; j < n && j <= i + d && arr[j] < arr[i]; j++ {
			dfs(arr, j, d)
			dp[i] = max(dp[i], dp[j] + 1)
		}
	}

	for i := 0 ;i < n; i++ {
		dfs(arr, i ,d)
	}

	for i := 0; i < n; i++ {
		ret = max(ret, dp[i])
	}
	return ret
}
// @lc code=end

