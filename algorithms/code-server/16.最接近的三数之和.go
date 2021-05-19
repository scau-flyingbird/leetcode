/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    d, ret := math.MaxInt32, 0
    for i := 0 ; i < n; i++ {
        for l, r := i + 1, n - 1; l < r; {
            sum := nums[i] + nums[l] + nums[r]
            dd := abs(sum, target)
            if dd < d {
                d = dd
                ret = sum
            }
            if sum == target {
                return sum
            }else if sum < target {
                l++
            }else {
                r--
            }
        }
    }

    return ret
}

func abs(a, b int) int {
    if a >= b {
        return a - b
    }else {
        return b - a
    }
}
// @lc code=end

