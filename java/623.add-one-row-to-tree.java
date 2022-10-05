/*
 * @lc app=leetcode id=623 lang=java
 *
 * [623] Add One Row to Tree
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
    public TreeNode addOneRow(TreeNode root, int val, int depth) {
        if(depth == 1) {
            TreeNode r = new TreeNode(val, root, null);
            return r;
        }

        this.iterate(root, 1, val, depth);

        return root;
    }

    private void iterate(TreeNode current, int ebene, int val, int depth) {
        if(current == null) return;

        if(ebene == (depth-1)) {
            TreeNode left = new TreeNode(val, current.left, null);
            TreeNode right = new TreeNode(val, null, current.right);

            current.left = left;
            current.right = right;
            return;
        }

        iterate(current.left, ebene+1, val, depth);
        iterate(current.right, ebene+1, val, depth);
    }
}
// @lc code=end

