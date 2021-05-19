/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func min (a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(findKthElement(nums1, nums2, total / 2 + 1))
	}else {
		return float64(findKthElement(nums1, nums2, total / 2) + findKthElement(nums1, nums2, total / 2 + 1)) / 2.0
	}
}
func findKthElement(nums1 []int, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2 + k - 1]
		}else if idx2 == len(nums2) {
			return nums1[idx1 + k - 1]
		}else if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}else {
			half := k / 2
			newIdx1 := min(idx1 + half, len(nums1)) - 1
			newIdx2 := min(idx2 + half, len(nums2)) - 1
			newVal1, newVal2 := nums1[newIdx1], nums2[newIdx2]
			if newVal1 <= newVal2 {
				k -= (newIdx1 - idx1 + 1)
				idx1 = newIdx1 + 1
			}else {
				k -= (newIdx2 - idx2 + 1)
				idx2 = newIdx2 + 1
			}
		}
	}
	return 0
}
// @lc code=end

