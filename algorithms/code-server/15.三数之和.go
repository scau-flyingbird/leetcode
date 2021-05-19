/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    ret := make([][]int, 0)

    //sort
    sort.Slice(nums, func (i, j int) bool {
        return nums[i] < nums[j]
    })

    // i... n - 1
    n := len(nums)
    preIdx := -1
    for i := 0; i < n; i++ {
        if preIdx != -1 && nums[i] == nums[preIdx] {
            continue
        }
        target := (-1) * nums[i]
        //l,r -> [i+1, n-1]
        for l, r := i + 1, n - 1; l < r; {
            sum := nums[l] + nums[r]
            if sum == target {
                ret = append(ret, []int{nums[i], nums[l], nums[r]})
                //unique
                v := nums[l]
                for ;l < r && nums[l] == v; l++ {}
            }else if sum < target {
                l++
            }else {
                r--
            }
        }
        preIdx = i
    }

    return ret
}
// @lc code=end

