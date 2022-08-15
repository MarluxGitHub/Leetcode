/*
 * @lc app=leetcode id=1379 lang=java
 *
 * [1379] Find a Corresponding Node of a Binary Tree in a Clone of That Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    public final TreeNode getTargetCopy(final TreeNode original, final TreeNode cloned, final TreeNode target) {
        return dfs(original, cloned, target);
    }

    private TreeNode dfs(TreeNode original, TreeNode cloned, TreeNode target) {
        if (original == target) {
            return cloned;
        }

        TreeNode left = null, right = null;
        if (original.left != null) {
            left = dfs(original.left, cloned.left, target);
        }
        if (original.right != null) {
            right = dfs(original.right, cloned.right, target);
        }

        return left != null ? left : right;
    }
}
// @lc code=end

