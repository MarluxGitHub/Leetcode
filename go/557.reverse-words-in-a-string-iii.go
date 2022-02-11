package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	sArray := strings.Split(s, " ")
	for i := 0; i < len(sArray); i++ {
		sArray[i] = reverseString([]byte(sArray[i]))
	}
	return strings.Join(sArray, " ")
}

func reverseString(s []byte) string {
    for i,j := 0, len(s) -1; i<j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
// @lc code=end

