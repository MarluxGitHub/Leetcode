/*
 * @lc app=leetcode id=540 lang=kotlin
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start
class Solution {
    fun singleNonDuplicate(nums: IntArray): Int = nums.reduce { a, b -> a xor b }
}
// @lc code=end

