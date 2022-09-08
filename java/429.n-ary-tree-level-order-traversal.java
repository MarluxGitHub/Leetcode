import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode id=429 lang=java
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    private List<List<Integer>> res;

    public List<List<Integer>> levelOrder(Node root) {
        res = new ArrayList<>();
        this.iterate(root, 0);

        return res;
    }

    private void iterate(Node current, int ebene) {
        if(current == null) return;
        if(res.size() <= ebene) res.add(new ArrayList<>());

        res.get(ebene).add(current.val);

        for (Node child : current.children) {
            this.iterate(child, ebene+1);
        }
    }
}
// @lc code=end

