/*
 * @lc app=leetcode id=1663 lang=golang
 *
 * [1663] Smallest String With A Given Numeric Value
 */

// @lc code=start
func getSmallestString(n int, k int) string {
	zCount := (k - n) / 25
	if k == 26 * zCount {
		return strings.Repeat("z", zCount)
	}

	aCount := n - zCount - 1
	x := k - 26 * zCount - aCount
	return strings.Repeat("a", aCount) + string('a' + x - 1) + strings.Repeat("z", zCount)
}

// @lc code=end

