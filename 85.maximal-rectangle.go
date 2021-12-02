/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    row := len(matrix)
    col := len(matrix[0])
    max := 0
    height := make([]int, col)
    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if matrix[i][j] == '1' {
                height[j]++
            } else {
                height[j] = 0
            }
        }
        max = maxMax(max, maxRectangle(height))
    }
    return max
}

func maxRectangle(height []int) int {
    if len(height) == 0 {
        return 0
    }
    stack := make([]int, 0)
    max := 0
    for i:=0; i<len(height); i++ {
        for len(stack) > 0 && height[stack[len(stack)-1]] >= height[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                max = maxMax(max, height[top]*i)
            } else {
                max = maxMax(max, height[top]*(i-stack[len(stack)-1]-1))
            }
        }
        stack = append(stack, i)
    }
    for len(stack) > 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if len(stack) == 0 {
            max = maxMax(max, height[top]*len(height))
        } else {
            max = maxMax(max, height[top]*(len(height)-stack[len(stack)-1]-1))
        }
    }
    return max
}

func maxMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

