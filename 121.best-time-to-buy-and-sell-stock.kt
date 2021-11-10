/*
 * @lc app=leetcode id=121 lang=kotlin
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
class Solution {
    fun maxProfit(prices: IntArray): Int {
        var maxProfit = 0
        var minPrice = Int.MAX_VALUE

        for (p in prices) {
            minPrice = minOf(minPrice, p)
            maxProfit = maxOf(maxProfit, p - minPrice)
        }

        return maxProfit
    }
}
// @lc code=end

