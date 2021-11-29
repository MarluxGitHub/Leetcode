/*
 * @lc app=leetcode id=106 lang=kotlin
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
    fun buildTree(inorder: IntArray, postorder: IntArray): TreeNode? {
        val inMap = inorder.mapIndexed{ index, value -> value to index}.toMap()
        return recursive(postorder.reversedArray(), inorder, inMap, 0, 0, inorder.size - 1)
    }

    fun recursive(revPost: IntArray,
                  inorder: IntArray,
                  inMap: Map<Int, Int>,
                  postIndex: Int,
                  inStart: Int,
                  inEnd:Int): TreeNode? {

        if (inStart > inEnd) return null
        val root = TreeNode(revPost[postIndex])
        val inIndex = inMap[root.`val`]
		// Note this is where ordering is flipped from the previous problem
        root.right = recursive(revPost, inorder, inMap, postIndex + 1, inIndex!! + 1, inEnd)
        root.left = recursive(revPost, inorder, inMap, postIndex + inEnd - inIndex!! + 1, inStart, inIndex - 1)
        return root
    }
}
// @lc code=end

