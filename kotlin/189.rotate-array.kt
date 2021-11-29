/*
 * @lc app=leetcode id=189 lang=kotlin
 *
 * [189] Rotate Array
 */

// @lc code=start
class Solution {
    fun rotate(nums: IntArray, k: Int): Unit {
        val k = k % nums.size
        reverse(nums, 0, nums.size - 1)
        reverse(nums, 0, k - 1)
        reverse(nums, k, nums.size - 1)
    }

    private fun reverse(nums: IntArray, start: Int, end: Int) {
        var i = start
        var j = end
        while (i < j) {
            val temp = nums[i]
            nums[i] = nums[j]
            nums[j] = temp
            i++
            j--
        }
    }
}

// @lc code=end

