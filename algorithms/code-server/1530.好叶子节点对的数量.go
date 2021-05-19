/*
 * @lc app=leetcode.cn id=1530 lang=golang
 *
 * [1530] 好叶子节点对的数量
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

func dfs(root *TreeNode, distance int, ret *int) []int {
	cnt := make([]int, distance)
	if root == nil {
		return cnt
	}

	lcnt := dfs(root.Left, distance, ret)
	rcnt := dfs(root.Right, distance, ret)

	for i := 0 ; i <= distance - 2; i++ {
		for j := 0; j <= distance - 2; j++ {
			if i + j + 2 > distance {
				continue
			}
			*ret = (*ret) + lcnt[i] * rcnt[j]
		}
	}

	for i := 0 ; i < distance - 1; i++ {
		cnt[i+1] = lcnt[i] + rcnt[i]
	}

	if root.Left == nil && root.Right == nil {
		cnt[0] = 1
	}
	return cnt
}

func countPairs(root *TreeNode, distance int) int {
	ret := 0
	dfs(root, distance, &ret)
	return ret
}
// @lc code=end

