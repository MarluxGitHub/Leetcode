package leetcode

/*
 * @lc app=leetcode id=342 lang=kotlin
 *
 * [342] Power of Four
 */

// @lc code=start
class Solution {
    fun isPowerOfFour(n: Int): Boolean {
        if (n == 0 || n == -2147483648) {
            return false
        }

        if (n == 1) {
            return true
        }

        if ((n.and(n-1)) == 0 && n.toString(2).length % 2 == 1) {
            return true
        }


        return false
    }
}
// @lc code=end

