/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	u := dividend > 0
	v := divisor > 0
	positive := (u && v) || (!u && !v)
	result := 0

	if !u {
		dividend = -dividend
	}

	if !v {
		divisor = -divisor
	}

	for dividend > 0 {
		q := 0
		dividend, q = helper(dividend, divisor)
		result += q
	}
	if positive {
		return result
	}

	return -result
}

func helper(dividend, divisor int) (int, int) {
	if dividend < divisor {
		return 0, 0
	}

	q := 1

	for (divisor << 1) < dividend {
		q = q << 1
		divisor = divisor << 1
	}

	return dividend - divisor, q
}
// @lc code=end

