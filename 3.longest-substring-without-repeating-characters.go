/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
	max := 0
	start := 0
	for i, v := range s {
		if j, ok := m[byte(v)]; ok && j >= start {
			start = j + 1
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		m[byte(v)] = i
	}
	return max
}
// @lc code=end

