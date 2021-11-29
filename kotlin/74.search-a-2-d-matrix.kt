/*
 * @lc app=leetcode id=74 lang=kotlin
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
class Solution {
    fun searchMatrix(matrix: Array<IntArray>, target: Int): Boolean {
        if (matrix.isEmpty()) return false
        var i = 0
        var j = matrix[0].size - 1
        while (i < matrix.size && j >= 0) {
            if (matrix[i][j] == target) return true
            if (matrix[i][j] > target) j--
            else i++
        }
        return false
    }
}
// @lc code=end

