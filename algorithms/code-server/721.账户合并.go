/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (this *UnionFind) find(x int) int {
	if x == this.parent[x] {
		return x
	} else {
		return this.find(this.parent[x])
	}
}

func (this *UnionFind) union(from int, to int) {
	pFrom := this.find(from)
	pTo := this.find(to)
	this.parent[pTo] = pFrom
}

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	emailToIndex := make(map[string]int, 0)

	p := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 1; j < len(accounts[i]); j++ {
			email := accounts[i][j]
			idx, ok := emailToIndex[email]
			if !ok {
				emailToIndex[email] = i
			} else {
				p.union(idx, i)
			}
		}
	}

	emails := make(map[int][]string, n)
	for email, idx := range emailToIndex {
		ufi := p.find(idx)
		emails[ufi] = append(emails[ufi], email)
	}

	ret := make([][]string, 0)
	for k, v := range emails {
		sort.Strings(v)
		acct := append([]string{accounts[k][0]}, v...)
		ret = append(ret, acct)
	}

	return ret
}

// @lc code=end

