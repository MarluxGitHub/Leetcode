package leetcode

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=902 lang=golang
 *
 * [902] Numbers At Most N Given Digit Set
 */

// @lc code=start
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
    sl, dl := len(s), float64(len(digits))
    var ans float64
    for i := 1; i < sl; i++ { ans += math.Pow(dl, float64(i)) }
    for i := 0; i < sl; i++ {
        prefix := false
        for _, d := range digits {
            if d[0] < s[i] {
                ans += math.Pow(dl, float64(sl - i - 1))
            } else if d[0] == s[i] {
                prefix = true
                break
            }
        }
        if !prefix { return int(ans) }
    }
    return int(ans + 1)
}
// @lc code=end

