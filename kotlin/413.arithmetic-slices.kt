/*
 * @lc app=leetcode id=413 lang=kotlin
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
class Solution {
    fun numberOfArithmeticSlices(nums: IntArray): Int {
        if(nums.size < 3) return 0

        var i = 0
        var j = 2
        var count = 0;

        var gap = nums[1] - nums[0];

        while(j < nums.size) {
            val currentGap = nums[j] - nums[j-1];

            if(currentGap.equals(gap)) {
                count += (j - i - 1);
            } else {
                i = j - 1;
                gap = currentGap;
            }
            j++
        }
        return count;
    }
}
// @lc code=end

