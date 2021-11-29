/*
 * @lc app=leetcode id=88 lang=kotlin
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
class Solution {
    fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
        var i = m + n - 1
        var j = n - 1
        var k = m - 1
        while (j >= 0) {
            nums1[i--] =  if (k >= 0 && nums1[k] > nums2[j]) nums1[k--] else  nums2[j--]
        }
    }
}
// @lc code=end

