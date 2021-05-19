/*
 * @lc app=leetcode.cn id=1289 lang=golang
 *
 * [1289] 下降路径最小和  II
 */

// @lc code=start
func minFallingPathSum(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	
	firSum, secSum, idx := 0, 0, -1
	for i := 0 ;i < m; i++ {
		ffSum, ssSum, iidx := math.MaxInt32, math.MaxInt32, -1
		for j := 0; j < n; j++ {
			var curSum int
			if j == idx {
				curSum = secSum + arr[i][j]
			}else {
				curSum = firSum + arr[i][j]
			}
			if curSum < ffSum {
				ssSum = ffSum
				ffSum = curSum
				iidx = j
			}else if curSum < ssSum {
				ssSum = curSum
			}
		}
		firSum = ffSum
		secSum = ssSum
		idx = iidx
	}

	return firSum
}
// @lc code=end

