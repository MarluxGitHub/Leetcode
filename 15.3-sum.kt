/*
 * @lc app=leetcode id=15 lang=kotlin
 *
 * [15] 3Sum
 */

// @lc code=start
class Solution {
    fun threeSum(nums: IntArray): List<List<Int>> {
        val sum:Int = 0
        nums.sort()

        var response:MutableList<List<Int>> = ArrayList()

        for(i in 0 .. nums.size - 3) {
            val k = sum - nums.get(i)
            var low = i + 1
            var high = nums.size - 1

            while(low < high) {
                val su = (nums.get(low) + nums.get(high))
                when  {
                    su < k -> low++
                    su > k -> high--
                    else -> {
                        val entry = listOf(nums.get(i), nums.get(low), nums.get(high))
                        if(!response.contains(entry)) response.add(entry)
                        low++
                        high--
                    }
                }
            }
        }
        return response
    }
}
// @lc code=end

