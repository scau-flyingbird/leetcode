/*
 * @lc app=leetcode.cn id=987 lang=golang
 *
 * [987] 二叉树的垂序遍历
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
	x int
	y int
	val int
}

func dfs(root *TreeNode, x int, y int, ret *[]Ret) {
	if root == nil {
		return 
	}
	*ret = append(*ret, Ret{x, y, root.Val})
	dfs(root.Left, x + 1, y - 1, ret)
	dfs(root.Right, x + 1, y + 1, ret)
}

func verticalTraversal(root *TreeNode) [][]int {
	ret := make([]Ret, 0)
	dfs(root, 0, 0, &ret)

	sort.Slice(ret, func (i, j int) bool{
		if ret[i].y < ret[j].y {
			return true
		}else if  ret[i].y > ret[j].y {
			return false
		}else {
			if ret[i].x < ret[j].x {
				return true
			}else if  ret[i].x > ret[j].x {
				return false
			}else {
				return ret[i].val < ret[j].val
			}
		}
	})

	arr := make([][]int, 0)

	tmp := make([]int, 0)
	for i := 0; i < len(ret); i++ {
		if i > 0 && ret[i].y != ret[i-1].y {
			arr = append(arr, tmp)
			tmp = make([]int, 0)
		}
		tmp = append(tmp, ret[i].val)
	}
	arr = append(arr, tmp)
	
	return arr
}
// @lc code=end

