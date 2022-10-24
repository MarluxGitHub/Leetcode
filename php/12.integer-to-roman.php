<?php
/*
 * @lc app=leetcode id=12 lang=php
 *
 * [12] Integer to Roman
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num)
    {
        $map = [
            1000 => 'M',
            900 => 'CM',
            500 => 'D',
            400 => 'CD',
            100 => 'C',
            90 => 'XC',
            50 => 'L',
            40 => 'XL',
            10 => 'X',
            9 => 'IX',
            5 => 'V',
            4 => 'IV',
            1 => 'I',
        ];

        $roman = '';
        foreach ($map as $int => $char) {
            while ($num >= $int) {
                $roman .= $char;
                $num -= $int;
            }
        }

        return $roman;
    }
}
// @lc code=end
