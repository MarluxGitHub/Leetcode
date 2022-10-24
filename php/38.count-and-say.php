<?php
/*
 * @lc app=leetcode id=38 lang=php
 *
 * [38] Count and Say
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer $n
     * @return String
     */
    function countAndSay($n)
    {
        if ($n == 1) {
            return  "1";
        }
        $s = $this->countAndSay($n - 1);
        $res = "";
        $count = 1;
        for ($i = 0; $i < strlen($s); $i++) {
            if ($i == strlen($s) || $s[$i + 1] != $s[$i]) {
                $res .= ((string) $count) . ((string)$s[$i]);
                $count = 1;
            } else {
                $count++;
            }
        }
        return $res;
    }
}
// @lc code=end
