/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start

//单词列表
var wl []string

//单词编号下标
var wm map[string]int

//单词最短路径滚动数组
var f []int

//构图
var e [][]int

var ret [][]string

//DFS查找最短路径
func dfs(st int, ed int, ans *[]string){
	*ans = append(*ans, wl[st])
	if st == ed {
		res := make([]string, len(*ans))
		copy(res, *ans)
		ret = append(ret, res)
	}

	for _, v := range e[st] {
		if f[st] == f[v] + 1 {
			dfs(v, ed, ans)
			*ans = (*ans)[:len(*ans) - 1]
		}
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	//把开始单词加入
	wordList = append(wordList, beginWord)

	//单词长度
	wlen := len(wordList)

	//初始化成员变量
	wl = wordList
	wm = make(map[string]int, wlen)
	f = make([]int, wlen)
	e = make([][]int, wlen)
	ret = make([][]string, 0)

	//把单词标号
	for i := 0; i < wlen; i++ {
		wm[wl[i]] = i
	}
	//特判返回
	endIdx, ok := wm[endWord]
	if !ok {
		return [][]string{}
	}

	//开始构图
	for i := 0; i < wlen; i++ {
		e[i] = make([]int, 0)
		str := []byte(wl[i])
		//遍历每个字母把每个字母改变成26位字母
		for j := 0; j < len(str); j++ {
			//当前字母
			curC := str[j]
			//开始str[j]替换成a-z
			for k := 0; k < 26; k++ {
				//排除掉原来字符
				if int(curC - 'a') != k {
					str[j] = byte('a' + k)
					//单词在列表中，则表示当前单词更改变一个字母成str
					if idx, ok := wm[string(str)]; ok {
						e[i] = append(e[i], idx)
					}
				}
			}
			str[j] = curC
		}
	}

	//BFS计算最短路径
	//初始化f[i]为最大值
	for i := 0; i < len(f); i++ {
		f[i] = math.MaxInt32
	}

	f[endIdx] = 0
	queue := []int{endIdx}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range e[u] {
			if f[u] + 1 < f[v] {
				f[v] = f[u] + 1
				queue = append(queue, v)
			}
		}
	}

	beginIdx := wm[beginWord]
	//fmt.Printf("least path length=%d", f[beginIdx])

	ans := make([]string, 0)
	dfs(beginIdx, endIdx, &ans)
	return ret
}
// @lc code=end

