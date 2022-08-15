/*
 * @lc app=leetcode id=991 lang=golang
 *
 * [991] Broken Calculator
 */

// @lc code=start
func brokenCalc(X int, Y int) int {
    ans := 0
    for Y > X {
        if Y % 2 == 0 {
            Y /= 2
        } else {
            Y += 1
        }
        ans++
    }
    return ans + X - Y
}
// @lc code=end

