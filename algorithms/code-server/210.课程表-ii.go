/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	e := make([][]int, numCourses)
	d := make([]int, numCourses)
	ret := make([]int, 0)

	for i := 0; i < len(e); i++ {
		e[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		e[prerequisites[i][1]] = append(e[prerequisites[i][1]], prerequisites[i][0])
		d[prerequisites[i][0]]++
	}

	//BFS
	queue := make([]int, 0)
	for i := 0; i < len(d); i++ {
		if d[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		ret = append(ret, u)

		for _,v := range e[u] {
			d[v]--
			if d[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(ret) < numCourses {
		return []int{}
	}

	return ret
}
// @lc code=end

