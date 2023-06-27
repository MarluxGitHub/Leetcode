/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
    var minDiff int = math.MaxInt32
    var stack []*TreeNode
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if prev != nil && node.Val-prev.Val < minDiff {
            minDiff = node.Val - prev.Val
        }
        prev = node
        node = node.Right
    }
    return minDiff
}
// @lc code=end

