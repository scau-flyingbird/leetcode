/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 */

// @lc code=start
func countTriplets(arr []int) int {
	ret := 0
	n := len(arr)
	if n == 0 {
		return 0
	}

	p := make([]int, n)
	p[0] = arr[0]

	for i := 1 ; i < n; i++ {
		p[i] = p[i-1]^arr[i]
	}

	for i := 0 ;i < n; i++ {
		for k := i + 1; k < n; k++ {
			//[i, k]
			var sum int
			if i == 0 {
				sum = p[k]
			}else {
				sum = p[k]^p[i-1]
			}
			if sum == 0 {
				ret += k - i
			}
		}
	}
	return ret
}
// @lc code=end

