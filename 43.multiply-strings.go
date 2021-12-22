/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
import (
	"bytes"
	"strconv"
)

func multiply(num1 string, num2 string) string {

	array1 := reverse(toIntArray(num1))
	array2 := reverse(toIntArray(num2))
	num := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			num[i+j] += array1[i] * array2[j]
		}
	}

	c := 0
	for i := 0; i < len(num); i++ {
		t := c + num[i]
		num[i] = t % 10
		c = t / 10
	}

	nonzero := len(num) - 1
	for ; nonzero >= 0 && num[nonzero] == 0; nonzero-- {

	}
	num = num[:nonzero+1]

	num = reverse(num)
	var buf bytes.Buffer
	for i := 0; i < len(num); i++ {
		buf.WriteString(strconv.Itoa(num[i]))
	}

	ans := buf.String()
	if len(ans) == 0 {
		return "0"
	} else {
		return ans
	}
}

func toIntArray(num string) []int {
	array := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		array[i], _ = strconv.Atoi(string(num[i]))
	}

	return array
}

func reverse(array []int) []int {
	i := 0
	j := len(array) - 1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}

	return array
}
// @lc code=end

