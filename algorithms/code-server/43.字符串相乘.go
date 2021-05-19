/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    n, m := len(num1), len(num2)
    ret := make([]int, n + m)
    
    for i := n - 1; i >= 0; i-- {
        x := int(num1[i] - '0')
        for j := m - 1; j >= 0; j-- {
            ret[i + j + 1] += x * int(num2[j] - '0')
        }
    }

    for i := m + n - 1; i > 0; i-- {
        ret[i - 1] += ret[i] / 10
        ret[i] = ret[i] % 10
    }

    idx := 0
    if ret[0] == 0 {
        idx = 1
    }
    
    str := ""
    for ; idx < m + n ; idx++ {
        str += strconv.Itoa(ret[idx])
    }
    return str
}
// @lc code=end

