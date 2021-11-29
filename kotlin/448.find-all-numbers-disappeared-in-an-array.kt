/*
 * @lc app=leetcode id=448 lang=kotlin
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
class Solution {
    fun findDisappearedNumbers(nums: IntArray): List<Int> {
        val result=MutableList(nums.size+1){it}
        nums.forEach{result[it]=0}
        result.removeAll{it==0}
        return result
    }
}
// @lc code=end

