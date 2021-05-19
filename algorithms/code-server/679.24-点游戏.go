/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */

// @lc code=start
const(
	TARGET = 24.0
	LIMIT = 1.0e-10
)

func calc(a, b float64, op int) float64 {
    switch op {
        case 0 : return a + b
        case 1 : return a - b
        case 2 : return a * b
        case 3 : return a / b
    }
    return -1
} 

func abs(a float64, b float64) bool{
    if a >= b {
        return a - b <= LIMIT
    }else {
        return b - a <= LIMIT
    }
}

func dfs(nums []float64) bool {
    if len(nums) == 1 {
        return abs(nums[0], TARGET)
    }

    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j {
                tmp := make([]float64, 0)
                for l := 0; l < len(nums); l++ {
                    if l == i || l == j {
                        continue
                    }
                    tmp = append(tmp, nums[l])
                }
                tmp = append(tmp, 0)
                for k := 0 ; k < 4; k++ {
                    ans := calc(nums[i], nums[j], k)
                    tmp[len(tmp) - 1] = ans
                    ret := dfs(tmp)
                    if ret {
                        return ret
                    }
                }
            }
        }
    }
    return false
}

func judgePoint24(nums []int) bool {
    arr := make([]float64, len(nums))
    for i := 0 ;i < len(nums); i++ {
        arr[i] = float64(nums[i])
    }

    return dfs(arr)
}
// @lc code=end

