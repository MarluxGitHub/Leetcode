/*
 * @lc app=leetcode id=1010 lang=golang
 *
 * [1010] Pairs of Songs With Total Durations Divisible by 60
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
    ans, cnt := 0, [60]int{}
	for _, t := range time {
		ans += cnt[(60 - t%60)%60]
		cnt[t%60]++
	}
	return ans
}
// @lc code=end

