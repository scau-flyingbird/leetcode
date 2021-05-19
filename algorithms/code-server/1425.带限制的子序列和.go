/*
 * @lc app=leetcode.cn id=1425 lang=golang
 *
 * [1425] 带限制的子序列和
 */

// @lc code=start
func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}

func constrainedSubsetSum(nums []int, k int) int {
    n := len(nums)
    queue := make([]int, 0)
    f := make([]int, n)
    ret := math.MinInt32

    for i := 0; i < n; i++ {
        //f[i] 表示选定i 最大的子序列和
        //[i-k, i - 1] 更新
        l := i - k - 1
        if len(queue) > 0 && queue[0] == l {
            queue = queue[1:]
        }

        f[i] = nums[i]
        if len(queue) > 0 {
            f[i] = max(f[i], f[queue[0]] + nums[i])
        }

        for len(queue) > 0 && f[queue[len(queue) - 1]] <= f[i] {
            queue = queue[:len(queue) - 1]
        }

        ret = max(ret, f[i])
        queue = append(queue, i)
    }
	return ret
}
// @lc code=end

