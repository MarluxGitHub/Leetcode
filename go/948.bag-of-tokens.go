package leetcode

import "sort"

/*
 * @lc app=leetcode id=948 lang=golang
 *
 * [948] Bag of Tokens
 */

// @lc code=start
func bagOfTokensScore(tokens []int, power int) int {
    	sort.Ints(tokens)
	ans, score := 0, 0
	i, j := 0, len(tokens)-1
	for i <= j {
		if power >= tokens[i] {
			power -= tokens[i]
			i++
			score++
			ans = max(ans, score)
		} else if score > 0 {
			power += tokens[j]
			j--
			score--
		} else {
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

