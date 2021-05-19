/*
 * @lc app=leetcode.cn id=1349 lang=golang
 *
 * [1349] 参加考试的最大学生数
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func checkOk(num int, seats []byte, m int) bool {
    pre := 0
	for i := 0; i < m; i++ {
		f := num >> i & 1
		if f == 1 && seats[i] == '#' {
			return false
		}
        if f == 1 && pre == 1 {
            return false
        }
        pre = f
	}
	return true
}

func checkOk2(s int, ss int, m int) bool {
	for i := 0 ;i < m ; i++ {
		b := s >> i & 1
        if b == 0 {
            continue
        }

        if i > 0 && ss >> (i - 1) & 1 == 1 ||  i + 1 < m && ss >> (i + 1) & 1 == 1 {
            return false
        }
	}
	return true
}

func maxStudents(seats [][]byte) int {
	n := len(seats)
	if n == 0 {
		return 0
	}
	m := len(seats[0])
	maxstate := 1 << m

	//初始化
	preX := make([]int, maxstate)
	preOk := make([]bool, maxstate)
	for i := 0; i < maxstate; i++ {
		preOk[i] = true
	}

	for i := 0; i < n; i++ {
		curX := make([]int, maxstate)
		curOk := make([]bool, maxstate)

		for s := 0; s < maxstate; s++ {

			//检查当前数字j是否符合要求
			curOk[s] = checkOk(s, seats[i], m)
			if !curOk[s] {
				continue
			}

			cnt := 0
			for i := 0; i < m; i++ {
				cnt += s >> i & 1
			}

			for ss := 0; ss < maxstate; ss++ {
				if !preOk[ss] {
					continue
				}
				if i == 0 || checkOk2(s, ss, m){
					curX[s] = max(curX[s], preX[ss] + cnt)
				}
			}
		}

		preX = curX
		preOk = curOk
	}

	maxRet := math.MinInt32

	for i := 0; i < maxstate; i++ {
		maxRet = max(maxRet, preX[i])
	}
	return maxRet
}
// @lc code=end

