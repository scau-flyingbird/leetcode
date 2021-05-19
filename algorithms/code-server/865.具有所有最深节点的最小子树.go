/*
 * @lc app=leetcode.cn id=865 lang=golang
 *
 * [865] 具有所有最深节点的最小子树
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
type Ret struct {
	dist int
	node *TreeNode
}

func dfs(root *TreeNode) Ret {
	if root == nil {
		return Ret{0, nil}
	}

	lret := dfs(root.Left)
	rret := dfs(root.Right)

	if lret.dist > rret.dist {
		return Ret{lret.dist + 1, lret.node}
	}else if lret.dist < rret.dist {
		return Ret{rret.dist + 1, rret.node}
	}else {
		return Ret{lret.dist + 1, root}
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	return dfs(root).node
}
// @lc code=end

