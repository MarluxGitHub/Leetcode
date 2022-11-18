/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
    	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	b := []byte(s)
	i, j := 0, len(b)-1
	for i < j {
		for i < len(b) && !vowels[b[i]] {
			i++
		}
		for j >= 0 && !vowels[b[j]] {
			j--
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b)
}
// @lc code=end

