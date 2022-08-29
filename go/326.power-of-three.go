package leetcode

import "math"

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	for i := 0; i < 31; i++ {
		if math.Pow(float64(3), float64(i)) == float64(n) {
			return true
		}
	}

	return false
}
// @lc code=end

