/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
    ret := letters[0]

	for _, v := range letters {
		if v > target {
			ret = v
			break
		}
	}

	return ret
}
// @lc code=end

