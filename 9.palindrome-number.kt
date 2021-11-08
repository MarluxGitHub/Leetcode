/*
 * @lc app=leetcode id=9 lang=kotlin
 *
 * [9] Palindrome Number
 */

// @lc code=start
class Solution {
    fun isPalindrome(x: Int): Boolean {
        return isStringPalindrome(x.toString())
    }

    fun isStringPalindrome(x: String): Boolean
    {
        for(i in 0 .. x.length/2) {
            if(x.get(i) != x.get(x.length-i-1)) return false
        }

        return true
    }
}
// @lc code=end

