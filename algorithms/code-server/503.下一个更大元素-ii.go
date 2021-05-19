/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	arr := make([]int, n * 2)
	copy(arr, nums)
	for i := n ;i < 2 * n; i++ {
		arr[i] = nums[i-n]
	}

	ret := make([]int, n)

	stack := make([]int, 0)
	for i := 2 * n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack) - 1]] <= arr[i] {
			stack = stack[:len(stack) - 1]
		}

		if i < n {
			if len(stack) > 0 && stack[len(stack) - 1] - i < n {
				ret[i] = arr[stack[len(stack) - 1]]
			}else {
				ret[i] = -1
			}
		}
		stack = append(stack, i)
	}

	return ret
}
// @lc code=end

