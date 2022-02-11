/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    start, end := 0, 0
    for i := 0; i < n; i++ {
        len1, len2 := expandAroundCenter(s, i, i), expandAroundCenter(s, i, i+1)
        len := max(len1, len2)
        if len > end - start {
            start = i - (len-1)/2
            end = i + len/2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

