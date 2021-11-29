/*
 * @lc app=leetcode id=260 lang=kotlin
 *
 * [260] Single Number III
 */

// @lc code=start
class Solution {
    fun singleNumber(nums: IntArray): IntArray {
        var mapper = mutableMapOf<Int, Int>()

        nums.forEach {
            mapper[it] = mapper.getOrDefault(it, 0) + 1
        }

        return (mapper.filter { predicate -> predicate.value == 1 }).keys.toIntArray()
    }
}
// @lc code=end

