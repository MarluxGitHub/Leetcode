/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	return x == res || x == res/10
}
// @lc code=end

