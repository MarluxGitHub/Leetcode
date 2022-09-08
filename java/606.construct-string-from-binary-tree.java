/*
 * @lc app=leetcode id=606 lang=java
 *
 * [606] Construct String from Binary Tree
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
    public String tree2str(TreeNode root) {
        return iterate(root);
    }

    public String iterate(TreeNode current) {
        String leftVal = "";
        String rightVal = "";
        if(current == null){
            return "";
        }

        if(current.left != null){
            leftVal = iterate(current.left);
        }

        if(current.right != null){
            rightVal = iterate(current.right);
        }


        if(!leftVal.equals("") || !rightVal.equals("")) leftVal = "("+leftVal+")";
        if(!rightVal.equals("")) rightVal = "(" + rightVal + ")";

        return current.val+ leftVal+rightVal;
    }
}
// @lc code=end

