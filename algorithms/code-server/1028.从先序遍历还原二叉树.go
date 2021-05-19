/*
 * @lc app=leetcode.cn id=1028 lang=golang
 *
 * [1028] 从先序遍历还原二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(pos *int, s string, depth int) *TreeNode {
	if *pos == len(s) {
		return nil
	}
	i := *pos
	for ; i < len(s) && s[i] == '-'; i++ {

	}
	if i - *pos != depth {
		return nil
	}

	val := 0
	for ; i < len(s) && s[i] != '-'; i++ {
		val = 10 * val + int(s[i] - '0')
	}
	*pos = i

	left := dfs(pos, s, depth+1)
	right := dfs(pos, s, depth+1)
	return &TreeNode{val, left, right}
}

func recoverFromPreorder(s string) *TreeNode {
	pos := 0
	return dfs(&pos, s, 0)
}
// @lc code=end

