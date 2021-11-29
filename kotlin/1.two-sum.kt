/*
 * @lc app=leetcode id=1 lang=kotlin
 *
 * [1] Two Sum
 */

// @lc code=start
class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val map = HashMap<Int, Int>()
        for (i in nums.indices) {
            if (map.containsKey(target - nums[i])) {
                return intArrayOf(map[target - nums[i]]!!, i)
            }
            map[nums[i]] = i
        }
        return intArrayOf(-1, -1)
    }
}
// @lc code=end

