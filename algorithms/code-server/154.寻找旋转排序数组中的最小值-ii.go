/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	low, hight := 0, len(nums) - 1
	for low < hight {
		mid := low + (hight - low)/2
		if nums[mid] < nums[hight]{
			hight = mid
		}else if nums[mid] > nums[hight] {
			low = mid + 1
		}else {
			hight--
		}
	}
	return nums[low]
}
// @lc code=end

