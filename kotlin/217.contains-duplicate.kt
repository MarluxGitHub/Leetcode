/*
 * @lc app=leetcode id=217 lang=kotlin
 *
 * [217] Contains Duplicate
 */

// @lc code=start
class Solution {
    fun containsDuplicate(nums: IntArray): Boolean {
        val set = HashSet<Int>()

        nums.forEach {
            if (set.contains(it)) {
                return true
            } else {
                set.add(it)
            }
         }

         return false
    }
}
// @lc code=end

