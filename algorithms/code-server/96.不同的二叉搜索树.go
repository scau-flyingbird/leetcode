/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	//G(n) = f(1) + f(2) + ...... + f(n)
	//f(n) = G(i-1)*G(n-i)
	//G(n) = G(0)*G(n-1)+ ...+G(n-1)*G(0)

	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}
// @lc code=end

