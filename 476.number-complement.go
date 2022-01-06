package leetcode

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	numStrBin := strconv.FormatInt(int64(num), 2)
	mask := math.Pow(2, float64(len(numStrBin))) - 1
	return num ^ int(mask)
}
// @lc code=end

