/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	idx := make([]int, len(nums))
	if len(nums) == 0 {
		return ret
	}

	for i := 0 ;i < len(nums); i++ {
		idx[i] = i
	}

	sort.Slice(idx, func (i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})

	i, j := 0, len(nums) - 1

	for i < j {
		if nums[idx[i]] + nums[idx[j]] > target {
			j--
		}else if nums[idx[i]] + nums[idx[j]] < target  {
			i++
		}else {
			ret[0], ret[1] = idx[i], idx[j]
			break
		}
	}
	return ret
}
// @lc code=end

