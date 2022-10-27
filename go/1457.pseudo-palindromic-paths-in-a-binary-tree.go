package leetcode

/*
 * @lc app=leetcode id=1457 lang=golang
 *
 * [1457] Pseudo-Palindromic Paths in a Binary Tree
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
func pseudoPalindromicPaths (root *TreeNode) int {
    var dfs func(*TreeNode, int) int
    dfs = func(node *TreeNode, mask int) int {
        if node == nil {
            return 0
        }
        mask ^= 1 << node.Val
        if node.Left == nil && node.Right == nil {
            if mask&(mask-1) == 0 {
                return 1
            }
            return 0
        }
        return dfs(node.Left, mask) + dfs(node.Right, mask)
    }
    return dfs(root, 0)
}
// @lc code=end

