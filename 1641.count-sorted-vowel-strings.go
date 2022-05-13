/*
 * @lc app=leetcode id=1641 lang=golang
 *
 * [1641] Count Sorted Vowel Strings
 */

// @lc code=start
func countVowelStrings(n int) int {
    if n == 1 {
        return 5
    }
    if n == 2 {
        return 15
    }
    if n == 3 {
        return 35
    }
    if n == 4 {
        return 70
    }

	n1h := (n - 1) * (n - 2)
	n2h := n1h * (n-3)

    five := n2h *(n - 4) / 24
    four := n2h * 5 / 6
    three := n1h * 5
    two := (n - 1) * 10
    return 5 + five + four + three + two
}
// @lc code=end

