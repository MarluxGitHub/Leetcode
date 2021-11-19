/*
 * @lc app=leetcode id=566 lang=kotlin
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
class Solution {
    fun matrixReshape(mat: Array<IntArray>, r: Int, c: Int): Array<IntArray> {

        if (mat.size * mat[0].size != r * c) {
            return mat
        }
        val result = Array(r) { IntArray(c) }
        var i = 0
        var j = 0
        for (row in mat) {
            for (item in row) {
                result[i][j] = item
                j++
                if (j == c) {
                    j = 0
                    i++
                }
            }
        }
        return result
    }
}
// @lc code=end

