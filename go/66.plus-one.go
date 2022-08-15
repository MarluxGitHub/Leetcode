/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {

	u := false
	for i := len(digits) - 1; i >= 0; i-- {
		u = false
		digits[i]++
		if(digits[i] == 10) {
			digits[i] = 0
			u = true
		}

		if(u == false ) {
			break
		}
	}

	if(u == true) {
		digits = append([]int{1}, digits...)
	}

	return digits
}
// @lc code=end

