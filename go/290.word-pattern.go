package leetcode

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {

	var dic = make(map[rune]string)
	var problemset = strings.Split(s, " ")

	if(len(pattern) != len(problemset)){
		return false
	}

	for pos, char := range pattern {
		if _, ok := dic[char]; ok {
			if problemset[pos] != dic[char]{
				return false
			}
		} else {
			if(containsMapValue(dic, problemset[pos])){
				return false
			}
			dic[char] = problemset[pos]
		}
	}

	return true
}

func containsMapValue(m map[rune]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}
// @lc code=end

