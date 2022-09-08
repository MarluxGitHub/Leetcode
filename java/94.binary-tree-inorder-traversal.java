import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode id=94 lang=java
 *
 * [94] Binary Tree Inorder Traversal
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
    private List<Integer> list = new ArrayList<>();

    public List<Integer> inorderTraversal(TreeNode root) {

        iterate(root);

        return list;
    }

    public void iterate(TreeNode current) {
        if (current != null) {
            iterate(current.left);
            list.add(current.val);
            iterate(current.right);
        }
    }
}
// @lc code=end

