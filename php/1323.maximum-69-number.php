<?
/*
 * @lc app=leetcode id=1323 lang=php
 *
 * [1323] Maximum 69 Number
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer $num
     * @return Integer
     */
    function maximum69Number($num)
    {
        return preg_replace('/6/', '9', (string) $num, 1);
    }
}
// @lc code=end
