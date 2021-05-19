/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)
	maxnum := 1 << n
	for i := 0; i < maxnum; i++ {
		tmp := make([]int, 0)
		for j := 0; j < n; j++ {
			f := i >> j & 1
			if f > 0 {
				tmp = append(tmp, nums[j])
			}
		}
		ret = append(ret, tmp)
	}

	return ret
}

// @lc code=end

