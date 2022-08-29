/*
 * @lc app=leetcode id=58 lang=kotlin
 *
 * [58] Length of Last Word
 */

// @lc code=start
class Solution {
    fun lengthOfLastWord(s: String): Int {
        return s.split(' ').for { s -> s.length }.filter { s -> s != 0 }
    }
}
// @lc code=end

