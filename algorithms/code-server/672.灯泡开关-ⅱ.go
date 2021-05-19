/*
 * @lc app=leetcode.cn id=672 lang=golang
 *
 * [672] 灯泡开关 Ⅱ
 */

// @lc code=start
func flipLights(n int, presses int) int {
	n = int(math.Min(float64(n), 3.0))
	if presses == 0 {
		return 1
	}else if presses == 1 {
		if n == 1 {
			return 2
		}else if n == 2 {
			return 3
		}else {
			return 4
		}
	}else if presses == 2 {
		if n == 1 {
			return 2
		}else if n == 2 {
			return 4
		}else {
			return 7
		}
	}else {
		if n == 1 {
			return 2
		}else if n == 2 {
			return 4
		}else {
			return 8
		}
	}
}
// @lc code=end

