package leetcode

import "strings"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)

	arr := strings.Split(s, " ")

	return len((arr[len(arr)-1]))
}
// @lc code=end

