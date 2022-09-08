/*
 * @lc app=leetcode id=814 lang=java
 *
 * [814] Binary Tree Pruning
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode pruneTree(TreeNode root) {
        iterate(root);

        if(root.val == 0 && root.left == null && root.right == null) return null;

        return root;
    }

    public boolean iterate(TreeNode current) {
        if(current == null) return false;
        boolean left = iterate(current.left);
        if(left == false) current.left = null;
        boolean right = iterate(current.right);
        if(right == false) current.right = null;
        if(current.val == 0 && !left && !right) return false;
        return true;
    }
}
// @lc code=end

