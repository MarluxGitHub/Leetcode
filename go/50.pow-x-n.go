/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
    	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n % 2 == 0 {
		return myPow(x * x, n / 2)
	}
	return x * myPow(x, n - 1)
}
// @lc code=end

