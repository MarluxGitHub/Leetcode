/*
 * @lc app=leetcode id=133 lang=javascript
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val, neighbors) {
 *    this.val = val === undefined ? 0 : val;
 *    this.neighbors = neighbors === undefined ? [] : neighbors;
 * };
 */

/**
 * @param {Node} node
 * @return {Node}
 */
var cloneGraph = function(node) {
    if (!node) return null;
  const map = new Map();
  const clone = new Node(node.val);
  map.set(node, clone);
  const queue = [node];
  while (queue.length) {
    const current = queue.shift();
    for (const neighbor of current.neighbors) {
      if (!map.has(neighbor)) {
        const cloneNeighbor = new Node(neighbor.val);
        map.set(neighbor, cloneNeighbor);
        map.get(current).neighbors.push(cloneNeighbor);
        queue.push(neighbor);
      } else {
        map.get(current).neighbors.push(map.get(neighbor));
      }
    }
  }
  return clone;
};
// @lc code=end

