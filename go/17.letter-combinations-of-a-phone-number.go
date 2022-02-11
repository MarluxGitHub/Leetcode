/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
    	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	result := []string{}
	for i := 0; i < len(digits); i++ {
		letters := m[digits[i]]
		if len(result) == 0 {
			for _, letter := range letters {
				result = append(result, letter)
			}
		} else {
			tmp := []string{}
			for _, letter := range letters {
				for _, s := range result {
					tmp = append(tmp, s+letter)
				}
			}
			result = tmp
		}
	}
	return result
}
// @lc code=end

