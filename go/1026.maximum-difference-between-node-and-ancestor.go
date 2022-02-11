/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
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
func maxAncestorDiff(root *TreeNode) int {
    if root == nil { return 0 }
    return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, min, max int) int {
	if root == nil { return max - min }
	if root.Val < min { min = root.Val }
	if root.Val > max { max = root.Val }
	return maxInt(dfs(root.Left, min, max), dfs(root.Right, min, max))
}

func maxInt(a, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end

