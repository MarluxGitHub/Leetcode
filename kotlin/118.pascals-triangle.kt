/*
 * @lc app=leetcode id=118 lang=kotlin
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
class Solution {
    fun generate(numRows: Int): List<List<Int>> {
        val result = mutableListOf<MutableList<Int>>()
        for (i in 0 until numRows) {
            result.add(mutableListOf())
            for (j in 0 until i + 1) {
                if (j == 0 || j == i) result[i].add(1)
                else {
                    result[i].add(result[i - 1][j - 1] + result[i - 1][j])
                }
            }
        }
        return result
    }
}
// @lc code=end

