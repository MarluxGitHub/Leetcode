/*
 * @lc app=leetcode id=1217 lang=golang
 *
 * [1217] Minimum Cost to Move Chips to The Same Position
 */

// @lc code=start
func minCostToMoveChips(position []int) int {
    odd, even := 0, 0
	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

