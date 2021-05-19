/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < n ; i++ {
		grid[0][i] += grid[0][i-1] 
	}

	for i := 1; i < m ; i++ {
		grid[i][0] += grid[i-1][0] 
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
// @lc code=end

