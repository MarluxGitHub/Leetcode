/*
 * @lc app=leetcode id=122 lang=kotlin
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
class Solution {

    fun maxProfit(prices: IntArray): Int {
        var maxProfit = 0

        for (i in 1 until prices.size) {
            if (prices[i] > prices[i-1]) {
                maxProfit += prices[i] - prices[i-1]
            }
        }

        return maxProfit
    }


}
// @lc code=end

