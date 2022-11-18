/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {
	b := []byte(s)
	i := 0
	for i < len(b)-1 {
		if b[i] == b[i+1] {
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

