/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var ret int

func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	ret = max(ret, left + right)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	ret = 0
	dfs(root)
	return ret
}
// @lc code=end

