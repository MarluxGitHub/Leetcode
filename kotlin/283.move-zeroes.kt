/*
 * @lc app=leetcode id=283 lang=kotlin
 *
 * [283] Move Zeroes
 */

// @lc code=start
class Solution {
    fun moveZeroes(nums: IntArray): Unit {
        var snowBallSize = 0;

        for(i in nums.indices) {
            if(nums[i] == 0) {
                snowBallSize++;
            } else {
                if (snowBallSize > 0) {
                    val t = nums[i];
                    nums[i] = 0;
                    nums[i-snowBallSize] = t;
                }
            }
        }
    }
}
// @lc code=end

