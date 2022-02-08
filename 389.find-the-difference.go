package leetcode

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
    sMap := make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}

	for _, v := range t {
		if(sMap[v] == 0) {
			return byte(v)
		} else {
			sMap[v]--
		}
	}

	for k, v := range sMap {
		if(v > 0) {
			return byte(k)
		}
	}

	return 0
}
// @lc code=end

