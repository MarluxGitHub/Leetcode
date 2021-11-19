/*
 * @lc app=leetcode id=668 lang=kotlin
 *
 * [668] Kth Smallest Number in Multiplication Table
 */

// @lc code=start
class Solution {
    fun findKthNumber(m: Int, n: Int, k: Int): Int {
        var l = 1
        var h = m*n
        var mid = l+((h-l)/2)

        while(l<h){
            var p = 0
            var p0 = 0

            for(j in 0 until m){
                p += minOf(mid/(j+1),n)
                p0 += minOf((mid-1)/(j+1),n)
            }

            when{
                k > p0 && k <= p -> return mid
                p0 >= k -> h = mid
                p < k -> l = mid+1
            }

            mid = l+((h-l)/2)
        }

        return mid
    }

}
// @lc code=end

