/*
 * @lc app=leetcode id=62 lang=kotlin
 *
 * [62] Unique Paths
 */

// @lc code=start
class Solution {
    fun uniquePaths(m: Int, n: Int): Int {

        var line = IntArray(n){1}

        for(i in 1 until m){
            for(j in 1 until n){
                line[j] += line[j-1]
            }
        }

        return line[n-1]
    }
}
// @lc code=end

