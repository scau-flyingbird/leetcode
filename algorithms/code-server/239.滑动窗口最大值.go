/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
    queue := make([]int, 0)

    push := func (i int) {
        for len(queue) > 0 && nums[queue[len(queue) - 1]] <= nums[i] {
            queue = queue[:len(queue) - 1]
        }
        queue = append(queue, i)
    }

    for i := 0; i < k; i++ {
        push(i)
    }

    ans := []int{nums[queue[0]]}
    for i := k; i < len(nums); i++ {
        push(i)
        for queue[0] <= i - k {
            queue = queue[1:]
        }
        ans = append(ans, nums[queue[0]])
    }
    return ans
}
// @lc code=end

