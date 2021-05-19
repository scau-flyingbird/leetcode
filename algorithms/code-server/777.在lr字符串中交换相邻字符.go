/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
func canTransform(start string, end string) bool {
	N := len(start)
	i , j := -1, -1
	for ;i < N && j < N; {
		
		if i < N {
			i++
		}
		if j < N {
			j++
		}

		for ; i < N && start[i] == 'X'; i++ {

		}
		for ; j < N && end[j] == 'X'; j++ {

		} 

		if i >= N && j < N || i < N && j >= N {
			return false
		}

		if i < N && j < N {
			if start[i] != end[j] || start[i] == 'L' && i < j || start[i] == 'R' && i > j {
				return false
			}
		}
	}

	return true
}
// @lc code=end

