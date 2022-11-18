<?
/*
 * @lc app=leetcode id=523 lang=php
 *
 * [523] Continuous Subarray Sum
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    function checkSubarraySum($nums, $k)
    {
        $sum = 0;
        $c =  count($nums);
        $map = [0 => -1];

        for ($i = 0; $i < $c; $i++) {
            $sum = ($sum + $nums[$i]) % $k;
            if ($sum == 0 && $i > 0) return true;

            if (!isset($map[$sum])) {
                $map[$sum] = $i;
            } else {
                if ($i - $map[$sum] > 1) return true;
            }
        }

        return false;
    }
}
