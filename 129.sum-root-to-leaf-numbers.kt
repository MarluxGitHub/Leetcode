/*
 * @lc app=leetcode id=129 lang=kotlin
 *
 * [129] Sum Root to Leaf Numbers
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
    var result = 0

    fun sumNumbers(root: TreeNode?): Int {
        addLeavesReq(root, result)
        return result
    }

    fun addLeavesReq(node: TreeNode?, current: Int) {
        if(node == null) return
        val c = current * 10 + node.`val`
        if(node.left == null && node.right == null) result += c
        addLeavesReq(node.left, c)
        addLeavesReq(node.right, c)
    }
}


// @lc code=end

