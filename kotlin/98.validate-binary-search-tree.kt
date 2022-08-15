/*
 * @lc app=leetcode id=98 lang=kotlin
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
/**
 * Example:
 * var ti = TreeNode(5)
 * var v = ti.`val`
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    fun isValidBST(root: TreeNode?): Boolean {
        return isValidBST(root, Long.MIN_VALUE, Long.MAX_VALUE)
    }

    fun isValidBST(root: TreeNode?, min: Long, max: Long): Boolean {
        if (root == null) return true
        if (root.`val`.toLong() <= min || root.`val`.toLong() >= max) return false
        return isValidBST(root.left, min.toLong(), root.`val`.toLong()) && isValidBST(root.right, root.`val`.toLong(), max)
    }
}
// @lc code=end

