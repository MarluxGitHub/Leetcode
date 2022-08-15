/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
    	if numRows == 1 {
		return s
	}
	n := len(s)
	if n <= numRows {
		return s
	}
	rows := make([][]byte, numRows)
	row, step := 0, 1
	for i := 0; i < n; i++ {
		rows[row] = append(rows[row], s[i])
		if row == 0 {
			step = 1
		}
		if row == numRows-1 {
			step = -1
		}
		row += step
	}
	var res string
	for _, row := range rows {
		res += string(row)
	}
	return res
}
// @lc code=end

