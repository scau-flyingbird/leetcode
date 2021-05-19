/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])

	step := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx := 0
	ret := make([]int, n * m)

	x, y, d := 0, -1, 0
	for i := 0; n > 0 && m > 0; i++ {
		if i % 2 == 0 {
			for k := 0; k < m; k++ {
				x += step[d][0]
				y += step[d][1]
				ret[idx] = matrix[x][y]
				idx++
			}
			n--
		}else {
			for k := 0; k < n; k++ {
				x += step[d][0]
				y += step[d][1]
				ret[idx] = matrix[x][y]
				idx++
			}
			m--
		}
		d = (d + 1) % 4
	}
	return ret
}
// @lc code=end

