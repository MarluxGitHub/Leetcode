/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
    	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	res := ""
	count := 1
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || s[i] != s[i+1] {
			res += strconv.Itoa(count) + string(s[i])
			count = 1
		} else {
			count++
		}
	}
	return res
}
// @lc code=end

