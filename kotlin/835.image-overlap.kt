/*
 * @lc app=leetcode id=835 lang=kotlin
 *
 * [835] Image Overlap
 */

// @lc code=start
class Solution {
    fun largestOverlap(img1: Array<IntArray>, img2: Array<IntArray>): Int {
        val n = img1.size
        val list1 = mutableListOf<Pair<Int, Int>>()
        val list2 = mutableListOf<Pair<Int, Int>>()
        for (i in 0 until n) {
            for (j in 0 until n) {
                if (img1[i][j] == 1) {
                    list1.add(Pair(i, j))
                }
                if (img2[i][j] == 1) {
                    list2.add(Pair(i, j))
                }
            }
        }
        val map = mutableMapOf<Pair<Int, Int>, Int>()
        for (p1 in list1) {
            for (p2 in list2) {
                val p = Pair(p1.first - p2.first, p1.second - p2.second)
                map[p] = map.getOrDefault(p, 0) + 1
            }
        }
        return map.values.max() ?: 0
    }
}
// @lc code=end

