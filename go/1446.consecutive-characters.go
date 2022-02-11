package leetcode

/*
 * @lc app=leetcode id=1446 lang=golang
 *
 * [1446] Consecutive Characters
 */

// @lc code=start
func maxPower(s string) int {
    power := 0
	maxPower := 0
	current := rune(s[0])

	for i := 0; i < len(s); i++ {
		if(current != rune(s[i])) {
			current = rune(s[i])
			if(power > maxPower) {
				maxPower = power
			}
			power = 1
		} else {
			power++
		}
	}

	if(power > maxPower) {
		maxPower = power
	}

	return maxPower
}
// @lc code=end

