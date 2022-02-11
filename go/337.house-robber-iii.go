/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
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
func rob(root *TreeNode) int {
    	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		return []int{
			node.Val + left[1] + right[1],
			max(left[0], left[1]) + max(right[0], right[1]),
		}
	}
	return max(dfs(root)[0], dfs(root)[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	return max(left+right, root.Val+left+right)
}
// @lc code=end

