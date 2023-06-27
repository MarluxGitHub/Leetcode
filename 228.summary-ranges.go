/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	ret := make([]string, 0)

	if len(nums) == 0 {
		return ret
	}

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == end + 1 {
			end = nums[i]
		} else {
			ret = append(ret, rangeToString(start, end))
			start, end = nums[i], nums[i]
		}
	}

	ret = append(ret, rangeToString(start, end))

	return ret
}

func rangeToString(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}

	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}
// @lc code=end

