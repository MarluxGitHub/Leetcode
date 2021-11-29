/*
 * @lc app=leetcode id=350 lang=kotlin
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
class Solution {
    fun intersect(nums1: IntArray, nums2: IntArray): IntArray {
        val map = mutableMapOf<Int, Int>()
        for (i in nums1) {
            map[i] = map.getOrDefault(i, 0) + 1
        }
        val res = mutableListOf<Int>()

        nums2.forEach {
            if(map.containsKey(it)){
                res.add(it)
                map[it] = map[it]!! - 1
                if(map[it] == 0) map.remove(it)
            }
         }

        return res.toIntArray()
    }
}
// @lc code=end

