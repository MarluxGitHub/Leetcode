/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && dict[s[i]] < dict[s[i+1]] {
			result -= dict[s[i]]
		} else {
			result += dict[s[i]]
		}
	}

	return result
}
// @lc code=end

