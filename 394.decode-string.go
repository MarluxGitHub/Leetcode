/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	ss := ""
	i := 0
	for i < len(s) {
		if !(s[i] >= '0' && s[i] <= '9') { // letter
			ss += string(s[i])
			i++
			continue
		}
		j := i
		for s[j] >= '0' && s[j] <= '9' {
			j++
		}
		num, _ := strconv.Atoi(s[i:j])
		k := j + 1
		count := 1
		for count > 0 {
			if s[k] == '[' {
				count++
			} else if s[k] == ']' {
				count--
			}
			k++
		}
		dec := decodeString(s[j+1 : k-1])
		ss += strings.Repeat(dec, num)
		i = k
	}
	return ss
}
// @lc code=end

