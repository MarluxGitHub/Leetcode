<?

/*
 * @lc app=leetcode id=223 lang=php
 *
 * [223] Rectangle Area
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer $ax1
     * @param Integer $ay1
     * @param Integer $ax2
     * @param Integer $ay2
     * @param Integer $bx1
     * @param Integer $by1
     * @param Integer $bx2
     * @param Integer $by2
     * @return Integer
     */
    function computeArea($ax1, $ay1, $ax2, $ay2, $bx1, $by1, $bx2, $by2)
    {
        $x1 = max($ax1, $bx1);
        $y1 = max($ay1, $by1);
        $x2 = min($ax2, $bx2);
        $y2 = min($ay2, $by2);

        $area1 = ($ax2 - $ax1) * ($ay2 - $ay1);
        $area2 = ($bx2 - $bx1) * ($by2 - $by1);
        $area3 = max($x2 - $x1, 0) * max($y2 - $y1, 0);

        return $area1 + $area2 - $area3;
    }
}
// @lc code=end
