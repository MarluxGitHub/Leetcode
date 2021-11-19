/*
 * @lc app=leetcode id=36 lang=kotlin
 *
 * [36] Valid Sudoku
 */

// @lc code=start
class Solution {
    fun isValidSudoku(board: Array<CharArray>): Boolean {
        val row = Array(9) { BooleanArray(9) }
        val col = Array(9) { BooleanArray(9) }
        val block = Array(9) { BooleanArray(9) }
        for (i in 0 until 9) {
            for (j in 0 until 9) {
                val c = board[i][j]
                if (c == '.') {
                    continue
                }
                val num = c - '1'
                if (row[i][num] || col[j][num] || block[i / 3 * 3 + j / 3][num]) {
                    return false
                }
                row[i][num] = true
                col[j][num] = true
                block[i / 3 * 3 + j / 3][num] = true
            }
        }
        return true
    }
}
// @lc code=end

