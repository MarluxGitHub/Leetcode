

import kotlin.time.seconds/*
 * @lc app=leetcode id=387 lang=kotlin
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
class Solution {
    fun firstUniqChar(s: String): Int {
        val map = HashMap<Char, Int>()

        for (i in s.indices) {
            map[s[i]] = map.getOrDefault(s[i], 0) + 1
        }
        for (i in s.indices) {
            if (map[s[i]] == 1) {
                return i
            }
        }
        return -1
    }
}
// @lc code=end

