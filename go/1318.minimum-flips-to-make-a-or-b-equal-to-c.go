package leetcode

/*
 * @lc app=leetcode id=1318 lang=golang
 *
 * [1318] Minimum Flips to Make a OR b Equal to c
 */

// @lc code=start
func minFlips(a int, b int, c int) int {
    res := 0
	for a > 0 || b > 0 || c > 0 {
		if c&1 == 1 {
			if a&1 == 0 && b&1 == 0 {
				res += 1
			}
		} else {
			if a&1 == 1 {
				res += 1
			}
			if b&1 == 1 {
				res += 1
			}
		}
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return res
}
// @lc code=end

