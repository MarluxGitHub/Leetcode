/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
    	if len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordNum := len(words)
	sLen := len(s)

	res := make([]int, 0)
	for i := 0; i <= sLen-wordNum*wordLen; i++ {
		if check(s[i:i+wordNum*wordLen], words) {
			res = append(res, i)
		}
	}
	return res
}

func check(s string, words []string) bool {
	wordLen := len(words[0])
	wordNum := len(words)
	m := make(map[string]int, wordNum)
	for _, word := range words {
		m[word]++
	}

	for i := 0; i < wordNum; i++ {
		if m[s[i*wordLen:(i+1)*wordLen]] == 0 {
			return false
		}
		m[s[i*wordLen:(i+1)*wordLen]]--
	}
	return true
}
// @lc code=end

