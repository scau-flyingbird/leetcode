/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	ans := make([]int, n)
	nmap := make(map[int]int, 0)

	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		nmap[nums2[i]] = i

		for len(stack) > 0 && stack[len(stack) - 1] <= nums2[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			ans[i] = -1
		}else {
			ans[i] = stack[len(stack) - 1]
		}

		stack = append(stack, nums2[i])
	}

	ret := make([]int, len(nums1))
	for i := 0 ; i < len(nums1); i++ {
		ret[i] = ans[nmap[nums1[i]]]
	}
	return ret
}
// @lc code=end

