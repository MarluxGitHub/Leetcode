/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
    	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	prev := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > prev {
			res++
			prev = points[i][1]
		}
	}
	return res
}
// @lc code=end

