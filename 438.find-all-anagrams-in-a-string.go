package leetcode

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	mismatchs := make([]int, 26)
	for i := 0; i < len(p); i++ {
		mismatchs[int(p[i]-'a')]++
	}

	for i := 0; i < len(p); i++ {
		mismatchs[int(s[i]-'a')]--
	}

	ans := []int{}
	if check(mismatchs) {
		ans = append(ans, 0)
	}

	for i := 1; i+len(p) <= len(s); i++ {
		end := i + len(p) - 1
		mismatchs[int(s[end]-'a')]--
		mismatchs[int(s[i-1]-'a')]++
		if check(mismatchs) {
			ans = append(ans, i)
		}
	}

	return ans
}

func check(mismatchs []int) bool {
	for _, mismatch := range mismatchs {
		if mismatch != 0 {
			return false
		}
	}

	return true
}
// @lc code=end

