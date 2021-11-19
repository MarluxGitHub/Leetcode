/*
 * @lc app=leetcode id=368 lang=kotlin
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start
class Solution {
    fun largestDivisibleSubset(nums: IntArray): List<Int> {
        if (nums.isEmpty()) {
            return emptyList()
        }
        nums.sort()
        val dp = IntArray(nums.size)
        val pre = IntArray(nums.size)
        var max = 0
        var maxIndex = 0
        for (i in nums.indices) {
            dp[i] = 1
            pre[i] = -1
            for (j in 0 until i) {
                if (nums[i] % nums[j] == 0 && dp[j] + 1 > dp[i]) {
                    dp[i] = dp[j] + 1
                    pre[i] = j
                }
            }
            if (dp[i] > max) {
                max = dp[i]
                maxIndex = i
            }
        }
        val res = ArrayList<Int>()
        while (maxIndex != -1) {
            res.add(nums[maxIndex])
            maxIndex = pre[maxIndex]
        }
        return res
    }
}
// @lc code=end

