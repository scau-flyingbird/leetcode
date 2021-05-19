/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
type Trie struct {
	next []*Trie
	isEnd bool
}

func Contruct() Trie{
	return Trie{next: make([]*Trie, 2)}
}

func (t *Trie) insert(x int) {
	h := t
	for i := 31; i >= 0; i-- {
		b := x >> i & 1
		if h.next[b] == nil {
			next := Trie{make([]*Trie, 2), false}
			h.next[b] = &next
		}
		h = h.next[b]
	}
	h.isEnd = true
}

func (t *Trie) queryMax(x int) int {
	ret, h := 0, t
	for i := 31; i >= 0; i-- {
		b := x >> i & 1
		b2 := 1 ^ b
		if h.next[b2] != nil {
			ret |= (1 << i)
			h = h.next[b2]
		}else {
			h = h.next[b]
		}
	}
	if !h.isEnd {
		return 0
	} else {
		return ret
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func findMaximumXOR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

    t := &Trie{make([]*Trie, 2), false}
	ret := 0
	t.insert(nums[0])
	for i := 1; i < len(nums); i++ {
		tmp := t.queryMax(nums[i])
		ret = max(ret, tmp)
		t.insert(nums[i])
	}
    return ret
}
// @lc code=end

