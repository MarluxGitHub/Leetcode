package leetcode

import "math/bits"

/*
 * @lc app=leetcode id=1009 lang=golang
 *
 * [1009] Complement of Base 10 Integer
 */

// @lc code=start
func bitwiseComplement(n int) int {
	if(n == 0) {return 1}
	return n ^ (1 << bits.Len32(uint32(n)) - 1)
}
// @lc code=end

