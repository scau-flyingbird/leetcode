/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
 func pathSum(root *TreeNode, sum int) int {
	cnt := 0
	path := make([]int, 1)

	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int)  {
		if root == nil {
			return
		}

		path = append(path, path[len(path)-1] + root.Val)

		n := len(path) - 1
		for i := n ; i > 0; i-- {
			if  path[n] - path[i-1] == sum {
				cnt++
			}
		}

		dfs(root.Left, sum)
		dfs(root.Right, sum)

		path = path[:len(path)-1]
	}

	dfs(root, sum)
	return cnt
}
// @lc code=end

