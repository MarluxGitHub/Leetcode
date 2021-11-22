/*
 * @lc app=leetcode id=1413 lang=kotlin
 *
 * [1413] Minimum Value to Get Positive Step by Step Sum
 */

// @lc code=start
class Solution {
    fun minStartValue(nums: IntArray): Int {
        var response = 1
        var sum = response

        for(num in nums) {
            sum += num
            if(sum < 1) {
                // compute diff if lesser then 1 und update response
                response += 1 - sum
                // add diff to sum to continue loop
                sum += 1 - sum
            }
        }

        return response
    }
}
// @lc code=end

