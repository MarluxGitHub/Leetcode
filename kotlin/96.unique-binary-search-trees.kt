/*
 * @lc app=leetcode id=96 lang=kotlin
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
class Solution {
    fun numTrees(n: Int): Long = (1..n).fold(1L){ catalan, i -> catalan * (i + n) / i } / (n + 1)
}
// @lc code=end

