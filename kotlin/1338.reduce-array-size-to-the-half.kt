/*
 * @lc app=leetcode id=1338 lang=kotlin
 *
 * [1338] Reduce Array Size to The Half
 */

// @lc code=start
class Solution {
    fun minSetSize(arr: IntArray): Int {
        val count = arr.size

        val map = HashMap<Int, Int>()
        for (i in arr) {
            map[i] = map.getOrDefault(i, 0) + 1
        }

        var current = 0
        var i = 0
        val sortedMap = map.entries.sortedByDescending { it.value }

        for (entry in sortedMap) {
            i++
            current = current + entry.value
            if(current >= count/2) {
                return i
            }
        }

        return -1
    }
}
// @lc code=end

