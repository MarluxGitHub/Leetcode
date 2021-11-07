/*
 * @lc app=leetcode id=13 lang=kotlin
 *
 * [13] Roman to Integer
 */

// @lc code=start
class Solution {
    fun romanToInt(s: String): Int {
        val map = hashMapOf<Char, Int>('I' to 1, 'V' to 5, 'X' to 10, 'L' to 50, 'C' to 100, 'D' to 500, 'M' to 1000)
        var result = 0
        for (i in s.indices) {
            if (i == s.lastIndex) {
                result += map[s[i]]!!
            } else if (map[s[i]]!! < map[s[i + 1]]!!) {
                result -= map[s[i]]!!
            } else {
                result += map[s[i]]!!
            }
        }
        return result
    }
}
// @lc code=end

