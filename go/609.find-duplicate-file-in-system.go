package leetcode

import "strings"

/*
 * @lc app=leetcode id=609 lang=golang
 *
 * [609] Find Duplicate File in System
 */

// @lc code=start
func findDuplicate(paths []string) [][]string {
    	m := make(map[string][]string)
	for _, path := range paths {
		tokens := strings.Split(path, " ")
		dir := tokens[0]
		for _, token := range tokens[1:] {
			i := strings.Index(token, "(")
			file := token[:i]
			content := token[i+1 : len(token)-1]
			m[content] = append(m[content], dir+"/"+file)
		}
	}
	res := [][]string{}
	for _, v := range m {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
// @lc code=end

