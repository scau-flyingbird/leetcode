/*
 * @lc app=leetcode.cn id=1125 lang=golang
 *
 * [1125] 最小的必要团队
 */

// @lc code=start
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    m := len(req_skills)
    n := len(people)
    maxstate := 1 << m

    f := make([]int64, maxstate)
    path := make([]int64, maxstate)
    sidx := make(map[string]int,0)

    for i := 0 ;i < m; i++ {
        sidx[req_skills[i]] = i
    }

    for j := 0 ;j < maxstate; j++ {
        f[j] = int64(n)
    }

    f[0] = 0
    for i := 1; i <= n; i++ {
        state := 0
        for _, ps := range people[i-1] {
            idx := sidx[ps]
            state |= (1 << idx)
        }

        //f[i-1][s] --> f[i][snew]
        for s := 0; s < maxstate; s++ {
            //选
            snew := s | state
            if f[s] + 1 < f[snew]{
                f[snew] = f[s] + 1
                path[snew] = path[s] | (1 << i)
            }
        }
    }

    ret := make([]int, 0)
    for i := 1; i <= n ; i++ {
        if (path[maxstate - 1] >> i) & 1 == 1 {
            ret = append(ret, i - 1)
        }
    }
    return ret
}
// @lc code=end

