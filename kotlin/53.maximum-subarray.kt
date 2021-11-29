/*
 * @lc app=leetcode id=53 lang=kotlin
 *
 * [53] Maximum Subarray
 */

// @lc code=start
class Solution {
    fun maxSubArray(nums: IntArray): Int {
        if (nums.isEmpty()) return 0
        var max = nums[0]
        var cur = nums[0]
        for (i in 1 until nums.size) {
            cur = maxOf(nums[i], cur + nums[i])
            max = maxOf(max, cur)
        }
        return max
    }
}
// @lc code=end

