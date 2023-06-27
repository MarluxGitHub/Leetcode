package leetcode

/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 { return true }
    if coordinates[0][0] == coordinates[1][0] {
        for _, c := range coordinates {
            if c[0] != coordinates[0][0] { return false }
        }
        return true
    }
    k := float64(coordinates[1][1] - coordinates[0][1]) / float64(coordinates[1][0] - coordinates[0][0])
    b := float64(coordinates[0][1]) - k * float64(coordinates[0][0])
    for _, c := range coordinates {
        if float64(c[1]) != k * float64(c[0]) + b { return false }
    }
    return true
}
// @lc code=end

