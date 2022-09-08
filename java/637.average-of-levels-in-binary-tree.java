import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode id=637 lang=java
 *
 * [637] Average of Levels in Binary Tree
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
    private List<List<Long>> res;

    public List<Double> averageOfLevels(TreeNode root) {
        res = new ArrayList<>();
        int ebene = 0;
        this.iterate(root, ebene);

        List<Double> response = new ArrayList<>();

        for(int i = 0; i < res.size(); i++) {
            long sum = 0;
            for(int j = 0; j < res.get(i).size(); j++)
            {
                sum += res.get(i).get(j);
            }
            double avg = ((double)sum) / res.get(i).size();
            response.add(avg);
        }

        return response;
    }

    private void iterate(TreeNode current, int ebene) {
        if(current == null) return;
        if(res.size() <= ebene) res.add(new ArrayList<>());
        res.get(ebene).add((long)current.val);
        this.iterate(current.left, ebene+1);
        this.iterate(current.right, ebene+1);
    }


}
// @lc code=end

