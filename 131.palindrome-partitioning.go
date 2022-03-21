/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
    	res := [][]string{}
	if len(s) == 0 {
		return res
	}
	dfs(s, 0, []string{}, &res)
	return res
}

func dfs(s string, start int, cur []string, res *[][]string) {
	if start == len(s) {
		*res = append(*res, append([]string{}, cur...))
		return
	}
	for i := start + 1; i <= len(s); i++ {
		if isPalindrome(s[start:i]) {
			dfs(s, i, append(cur, s[start:i]), res)
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
// @lc code=end

