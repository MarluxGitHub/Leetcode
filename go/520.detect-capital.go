package leetcode

import "unicode"

/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	if(len(word) < 2) { return true}
	allUpper := unicode.IsUpper(rune(word[0])) && unicode.IsUpper(rune(word[1]))

	for i := 1; i < len(word); i++ {
		if(unicode.IsUpper(rune(word[i])) != allUpper) {return false}
	}

	return true
}
// @lc code=end

