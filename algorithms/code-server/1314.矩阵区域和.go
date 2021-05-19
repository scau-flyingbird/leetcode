/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		sum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] + mat[i-1][j-1] - sum[i-1][j-1]
		}
		fmt.Printf("\n")
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			r1, c1 := max(i-k, 1), max(j-k, 1)
			r2, c2 := min(i+k, m), min(j+k, n)
			mat[i-1][j-1] = sum[r2][c2] - sum[r2][c1-1] - sum[r1-1][c2] + sum[r1-1][c1-1]
		}
	}
	return mat
}

// @lc code=end

