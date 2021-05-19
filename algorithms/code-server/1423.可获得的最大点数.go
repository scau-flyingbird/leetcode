/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}

func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    for i := 1; i < n; i++ {
        cardPoints[i] += cardPoints[i-1]
    }

    ret := 0
    l, r := k - 1, n
    for ; l >= -1; {
        var lsum, rsum int
        if l == -1 {
            lsum = 0
        }else {
            lsum = cardPoints[l]
        }

        if r == n {
            rsum = 0
        }else if r - 1 < 0{
            rsum = cardPoints[n-1]
        }else {
            rsum = cardPoints[n-1] - cardPoints[r-1]
        }

        ret = max(ret, lsum + rsum)
        l--
        r--
    }
    return ret
}
// @lc code=end

