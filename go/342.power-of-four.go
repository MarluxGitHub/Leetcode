package leetcode

import (
	"math/bits"
)

/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n&(n-1) == 0 && bits.TrailingZeros32(uint32(n)) % 2 == 0 {
		return true
	}

	return false
}
// @lc code=end

