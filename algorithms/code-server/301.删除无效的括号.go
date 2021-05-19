/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
type Pair struct {
	ch byte
	cnt int
}
var ret []string
var pairs []Pair
var maxlength int

func gen(ans *[]Pair) string {
	res := make([]byte, 0)
	for _, pair := range *ans {
		for j := 0; j < pair.cnt; j++ {
			res = append(res, pair.ch)
		}
	}
	return string(res)
}

func dfs(cur int, s int, l int, ans *[]Pair) {
	if cur == len(pairs) {
		if l == maxlength && s == 0 {
			ret = append(ret, gen(ans))
		}
        return
	}

	//不是括号
	if pairs[cur].ch != '(' && pairs[cur].ch != ')' {
		if l + pairs[cur].cnt > maxlength {
			return
		}
		*ans = append(*ans, pairs[cur])
		dfs(cur + 1, s, l + pairs[cur].cnt, ans)
		*ans = (*ans)[:len(*ans) - 1]
		return
	}

	//枚举(((  )))
	*ans = append(*ans, pairs[cur])
	for i := 0; i <= pairs[cur].cnt; i++ {
		if pairs[cur].ch == '(' {
			if i + l <= maxlength {
				(*ans)[cur].cnt = i
				dfs(cur + 1, s + i, l + i, ans)
			}
		}else {
			if s >= i && i + l <= maxlength {
				(*ans)[cur].cnt = i
				dfs(cur + 1, s - i, l + i, ans)
			}
		}
	}
	*ans = (*ans)[:len(*ans) - 1]
}
func removeInvalidParentheses(s string) []string {
	//初始化
	ret = make([]string, 0)
	str := []byte(s)
	pairs = make([]Pair, 0)

	cnts, cntd, cnt := 0, 0, 0
	i := 0
	for ; i < len(str); i++ {
		if i == 0 || str[i] == str[i-1]{
			cnt++
		}else {
			pairs = append(pairs, Pair{str[i-1], cnt})
			cnt = 1
		}

		if str[i] == '(' {
			cnts++
		}else if str[i] == ')'{
			if cnts > 0 {
				cnts--
			}else {
				cntd++
			}
		}
	}

	pairs = append(pairs, Pair{str[i-1], cnt})
	maxlength = len(str) - (cnts + cntd)

	ans := make([]Pair, 0)
	dfs(0, 0, 0, &ans)
	return ret
}
// @lc code=end

