/*
 * @lc app=leetcode id=35 lang=kotlin
 *
 * [35] Search Insert Position
 */

// @lc code=start
class Solution {
    fun searchInsert(nums: IntArray, target: Int): Int {
        var left = 0
        var right = nums.size - 1
        var mid = (left + right) / 2

        while (left <= right) {
            when {
                nums[mid] == target -> return mid
                nums[mid] > target -> right = mid - 1
                else -> left = mid + 1
            }
            mid = (left + right) / 2
        }

        return if(nums[mid] > target) mid else mid + 1
    }
}
// @lc code=end

