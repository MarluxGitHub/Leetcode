/*
 * @lc app=leetcode id=1161 lang=golang
 *
 * [1161] Maximum Level Sum of a Binary Tree
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
func maxLevelSum(root *TreeNode) int {
	var maxSum, maxLevel int = root.Val, 1
    var queue []*TreeNode
    queue = append(queue, root)
    level := 1
    for len(queue) > 0 {
        sum := 0
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[i]
            sum += node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if sum > maxSum {
            maxSum = sum
            maxLevel = level
        }
        queue = queue[size:]
        level++
    }
    return maxLevel
}
// @lc code=end

