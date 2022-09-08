package leetcode

/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return goodNodesDFS(root, root.Val)
}

func goodNodesDFS(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	if root.Val >= max {
		max = root.Val
		return 1 + goodNodesDFS(root.Left, max) + goodNodesDFS(root.Right, max)
	}
	return goodNodesDFS(root.Left, max) + goodNodesDFS(root.Right, max)
}
// @lc code=end
