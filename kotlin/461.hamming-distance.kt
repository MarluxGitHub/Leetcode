/*
 * @lc app=leetcode id=461 lang=kotlin
 *
 * [461] Hamming Distance
 */

// @lc code=start
class Solution {
    fun hammingDistance(x: Int, y: Int) = Integer.bitCount(x xor y)
}
// @lc code=end

