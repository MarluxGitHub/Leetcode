/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
    	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	var res int64
	for _, v := range s {
		if v < '0' || v > '9' {
			break
		}
		res = res*10 + int64(v-'0')
		if res > math.MaxInt32 {
			break
		}
	}
	res *= int64(sign)
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return int(res)
}
// @lc code=end

