package leetcode

/*
 * @lc app=leetcode id=1544 lang=golang
 *
 * [1544] Make The String Great
 */

// @lc code=start
func makeGood(s string) string {
	b := []byte(s)
	i := 0
	for i < len(b)-1 {
		if b[i] == b[i+1] || b[i] == b[i+1]-32 {
			b = append(b[:i], b[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return string(b)
}
// @lc code=end

