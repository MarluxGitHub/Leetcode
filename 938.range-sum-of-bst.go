/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
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
func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
	return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}


// @lc code=end

