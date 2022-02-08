package leetcode

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}
	var pos int

	for pos < len(gas) {
		if gas[pos] >= cost[pos] {
			endPoint := run(gas, cost, pos)
			if endPoint < pos {
				return -1
			} else if endPoint == pos {
				return pos
			}
			pos = endPoint + 1
		} else {
			pos++
		}
	}

	return -1
}

func run(gas, cost []int, pos int) int {
	tank := gas[pos] - cost[pos]
	endPoint := pos

	for tank >= 0 {
		endPoint++
		if endPoint == len(gas) {
			endPoint = 0
		}
		if endPoint == pos {
			break
		}
		tank += (gas[endPoint] - cost[endPoint])
	}

	return endPoint
}

func sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
// @lc code=end

