/*
 * @lc app=leetcode id=404 lang=kotlin
 *
 * [404] Sum of Left Leaves
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

    fun sumOfLeftLeaves(root: TreeNode?, isLeft:Boolean = false): Int {
        if(root == null) return 0
        if (root!!.left == null && root.right == null){
            return if(isLeft) root.`val` else 0
        }
        return sumOfLeftLeaves(root.left,true) + sumOfLeftLeaves(root.right,false)
    }
}
// @lc code=end

