/*
 * @lc app=leetcode id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
    maxAltitude, currentAltitude := 0, 0
    for _, g := range gain {
        currentAltitude += g
        if currentAltitude > maxAltitude {
            maxAltitude = currentAltitude
        }
    }
    return maxAltitude
}
// @lc code=end

