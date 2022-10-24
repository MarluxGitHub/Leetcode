<?php
/*
 * @lc app=leetcode id=1239 lang=php
 *
 * [1239] Maximum Length of a Concatenated String with Unique Characters
 */

// @lc code=start
class Solution
{

    /**
     * @param String[] $arr
     * @return Integer
     */
    function maxLength($arr)
    {
        $res = 0;
        $this->backtrack($arr, 0, '', $res);
        return $res;
    }

    function backtrack($arr, $index, $str, &$res)
    {
        if ($index == count($arr)) {
            $res = max($res, strlen($str));
            return;
        }
        if ($this->isUnique($str . $arr[$index])) {
            $this->backtrack($arr, $index + 1, $str . $arr[$index], $res);
        }
        $this->backtrack($arr, $index + 1, $str, $res);
    }

    function isUnique($str)
    {
        $arr = str_split($str);
        return count($arr) == count(array_unique($arr));
    }
}
// @lc code=end
