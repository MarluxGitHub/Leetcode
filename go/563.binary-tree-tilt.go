/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
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
func findTilt(root *TreeNode) int {
    	if root == nil {
		return 0
	}
	return findTilt(root.Left) + findTilt(root.Right) + abs(sum(root.Left) - sum(root.Right))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}
// @lc code=end

