/*
 * @lc app=leetcode id=704 lang=kotlin
 *
 * [704] Binary Search
 */

// @lc code=start
class Solution {
    fun search(nums: IntArray, target: Int): Int {
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

        return -1

    }
}
// @lc code=end

