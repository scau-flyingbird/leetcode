/*
 * @lc app=leetcode.cn id=877 lang=golang
 *
 * [877] 石子游戏
 */

// @lc code=start
func max(a,b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}
func stoneGame(piles []int) bool {
	n := len(piles)
	f := make([][]int, n)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
		f[i][i] = piles[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			//f[i][j]
			f[i][j] = max(piles[i] - f[i+1][j], piles[j] - f[i][j-1])
		}
	}
	return f[0][n-1] > 0
}
// @lc code=end

