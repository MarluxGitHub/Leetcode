/*
 * @lc app=leetcode id=235 lang=kotlin
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int = 0) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */

class Solution {
    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        val lca = root

        if (p!!.`val` < root!!.`val` && q!!.`val` < root!!.`val`) {
            return lowestCommonAncestor(root.left, p, q);
        }

        if (p!!.`val` > root!!.`val` && q!!.`val` > root!!.`val`) {
            return lowestCommonAncestor(root.right, p, q);
        }

        return lca
    }
}
// @lc code=end

