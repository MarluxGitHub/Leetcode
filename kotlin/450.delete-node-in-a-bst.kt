/*
 * @lc app=leetcode id=450 lang=kotlin
 *
 * [450] Delete Node in a BST
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
    fun deleteNode(root: TreeNode?, key: Int): TreeNode? {
        return root?.apply {
            if (`val` == key) {
                if (left == null) return right
                if (right == null) return left
                findSuccessor()?.also {
                    `val` = it.`val`
                    right = deleteNode(right, it.`val`)
                }
            }
            if (`val` < key) right = deleteNode(right, key)
            else left = deleteNode(left, key)
        }
    }

    private fun TreeNode.findSuccessor(): TreeNode? {
        var r = right
        while (r?.left != null) {
            r = r.left
        }
        return r
    }
}
// @lc code=end

