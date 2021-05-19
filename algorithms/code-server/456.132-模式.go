/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132 模式
 */

// @lc code=start
func min(a, b int) int{
    if a < b {
        return a
    }else {
        return b
    }
}

func find132pattern(nums []int) bool {
    //i < j < k. nums[i] < nums[k] < nums[j]
    n := len(nums)
    mins := make([]int, n)
    stack := make([]int, 0)

    for i := 0 ;i < n; i++ {
        // mins[i] 1..i最小前缀
        if i == 0 {
            mins[i] = nums[i]
        }else {
            mins[i] = min(mins[i-1], nums[i])
        }

        for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
            stack = stack[:len(stack) - 1]
        }

        if len(stack) > 0 && stack[len(stack) - 1] - 1 >= 0 && mins[stack[len(stack) - 1] - 1] < nums[i] {
            return true
        }

        stack = append(stack, i)
    }

    return false
}
// @lc code=end

