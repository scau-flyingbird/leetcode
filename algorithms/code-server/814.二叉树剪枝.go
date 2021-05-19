/*
 * @lc app=leetcode.cn id=814 lang=golang
 *
 * [814] 二叉树剪枝
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

func dfs(root *TreeNode) bool {
	if root == nil {
		return false
	}

	lret := dfs(root.Left)
	if !lret {
		root.Left = nil
	}
	rret := dfs(root.Right)
	if !rret {
		root.Right = nil
	}

	if root.Val == 0 && !lret && !rret {
		return false
	}else {
		return true
	}
}

func pruneTree(root *TreeNode) *TreeNode {
	if !dfs(root) {
		return nil
	}else {
		return root
	}
}
// @lc code=end

