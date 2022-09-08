/*
 * @lc app=leetcode id=48 lang=kotlin
 *
 * [48] Rotate Image
 */

// @lc code=start
class Solution {
    fun rotate(matrix: Array<IntArray>): Unit {
        val n = matrix.size
        for (i in 0 until n / 2) {
            for (j in i until n - i - 1) {
                val temp = matrix[i][j]
                matrix[i][j] = matrix[n - 1 - j][i]
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
                matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
                matrix[j][n - 1 - i] = temp
            }
        }
    }
}
// @lc code=end

